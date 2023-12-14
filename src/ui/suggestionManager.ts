// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import { Suggestion, SuggestionBlob } from "../runtime/model.js";
import { getSuggestions } from "../runtime/runtime.js";
import { ISTerm } from "../isterm/pty.js";
import { renderBox, truncateText, truncateMultilineText } from "./utils.js";
import ansi from "ansi-escapes";
import chalk from "chalk";
import { parseKeystroke } from "../utils/ansi.js";

const maxSuggestions = 5;
const suggestionWidth = 40;
const descriptionWidth = 30;
const descriptionHeight = 5;
const borderWidth = 2;
const activeSuggestionBackgroundColor = "#7D56F4";
export const MAX_LINES = borderWidth + Math.max(maxSuggestions, descriptionHeight);
type SuggestionsSequence = {
  data: string;
  rows: number;
};

export class SuggestionManager {
  #term: ISTerm;
  #command: string;
  #activeSuggestionIdx: number;
  #suggestBlob?: SuggestionBlob;

  constructor(terminal: ISTerm) {
    this.#term = terminal;
    this.#suggestBlob = { suggestions: [] };
    this.#command = "";
    this.#activeSuggestionIdx = 0;
  }

  private async _loadSuggestions(): Promise<void> {
    const commandText = this.#term.getCommandState().commandText;
    if (!commandText) {
      this.#suggestBlob = undefined;
      return;
    }
    if (commandText == this.#command) {
      return;
    }
    this.#command = commandText;
    const suggestionBlob = await getSuggestions(commandText);
    this.#suggestBlob = suggestionBlob;
  }

  private _renderArgumentDescription(description: string | undefined, x: number) {
    if (!description) return "";
    return renderBox([truncateText(description, descriptionWidth - borderWidth)], descriptionWidth, x);
  }

  private _renderDescription(description: string | undefined, x: number) {
    if (!description) return "";
    return renderBox(truncateMultilineText(description, descriptionWidth - borderWidth, descriptionHeight), descriptionWidth, x);
  }

  private _descriptionRows(description: string | undefined) {
    if (!description) return 0;
    return truncateMultilineText(description, descriptionWidth - borderWidth, descriptionHeight).length;
  }

  private _renderSuggestions(suggestions: Suggestion[], activeSuggestionIdx: number, x: number) {
    return renderBox(
      suggestions.map((suggestion, idx) => {
        const suggestionText = `${suggestion.icon} ${suggestion.name}`.padEnd(suggestionWidth - borderWidth, " ");
        const truncatedSuggestion = truncateText(suggestionText, suggestionWidth - 2);
        return idx == activeSuggestionIdx ? chalk.bgHex(activeSuggestionBackgroundColor)(truncatedSuggestion) : truncatedSuggestion;
      }),
      suggestionWidth,
      x,
    );
  }

  async render(remainingLines: number): Promise<SuggestionsSequence> {
    await this._loadSuggestions();
    if (!this.#suggestBlob) return { data: "", rows: 0 };
    const { suggestions, argumentDescription } = this.#suggestBlob;

    const page = Math.min(Math.floor(this.#activeSuggestionIdx / maxSuggestions) + 1, Math.floor(suggestions.length / maxSuggestions) + 1);
    const pagedSuggestions = suggestions.filter((_, idx) => idx < page * maxSuggestions && idx >= (page - 1) * maxSuggestions);
    const activePagedSuggestionIndex = this.#activeSuggestionIdx % maxSuggestions;
    const activeDescription = pagedSuggestions.at(activePagedSuggestionIndex)?.description || argumentDescription || "";

    const wrappedPadding = this.#term.getCursorState().cursorX % this.#term.cols;
    const maxPadding = activeDescription.length !== 0 ? this.#term.cols - suggestionWidth - descriptionWidth : this.#term.cols - suggestionWidth;
    const swapDescription = wrappedPadding > maxPadding && activeDescription.length !== 0;
    const swappedPadding = swapDescription ? Math.max(wrappedPadding - descriptionWidth, 0) : wrappedPadding;
    const clampedLeftPadding = Math.min(Math.min(wrappedPadding, swappedPadding), maxPadding);

    if (suggestions.length <= this.#activeSuggestionIdx) {
      this.#activeSuggestionIdx = Math.max(suggestions.length - 1, 0);
    }

    if (pagedSuggestions.length == 0) {
      if (argumentDescription != null) {
        return {
          data:
            ansi.cursorHide +
            ansi.cursorUp(2) +
            ansi.cursorForward(clampedLeftPadding) +
            this._renderArgumentDescription(argumentDescription, clampedLeftPadding),
          rows: 3,
        };
      }
      return { data: "", rows: 0 };
    }

    const suggestionRowsUsed = pagedSuggestions.length + borderWidth;
    let descriptionRowsUsed = this._descriptionRows(activeDescription) + borderWidth;
    let rows = Math.max(descriptionRowsUsed, suggestionRowsUsed);
    if (rows <= remainingLines) {
      descriptionRowsUsed = suggestionRowsUsed;
      rows = suggestionRowsUsed;
    }

    const descriptionUI =
      ansi.cursorUp(descriptionRowsUsed - 1) +
      (swapDescription
        ? this._renderDescription(activeDescription, clampedLeftPadding)
        : this._renderDescription(activeDescription, clampedLeftPadding + suggestionWidth)) +
      ansi.cursorDown(descriptionRowsUsed - 1);
    const suggestionUI =
      ansi.cursorUp(suggestionRowsUsed - 1) +
      (swapDescription
        ? this._renderSuggestions(pagedSuggestions, activePagedSuggestionIndex, clampedLeftPadding + descriptionWidth)
        : this._renderSuggestions(pagedSuggestions, activePagedSuggestionIndex, clampedLeftPadding)) +
      ansi.cursorDown(suggestionRowsUsed - 1);

    const ui = swapDescription ? descriptionUI + suggestionUI : suggestionUI + descriptionUI;
    return {
      data: ansi.cursorHide + ansi.cursorForward(clampedLeftPadding) + ui + ansi.cursorShow,
      rows,
    };
  }

  update(input: Buffer): "handled" | "fully-handled" | false {
    const keyStroke = parseKeystroke(input);
    if (keyStroke == null) return false;
    if (keyStroke == "esc") {
      this.#suggestBlob = undefined;
    } else if (keyStroke == "up") {
      this.#activeSuggestionIdx = Math.max(0, this.#activeSuggestionIdx - 1);
    } else if (keyStroke == "down") {
      this.#activeSuggestionIdx = Math.min(this.#activeSuggestionIdx + 1, (this.#suggestBlob?.suggestions.length ?? 1) - 1);
    } else if (keyStroke == "tab") {
      const removals = "\u007F".repeat(this.#suggestBlob?.charactersToDrop ?? 0);
      const chars = this.#suggestBlob?.suggestions.at(this.#activeSuggestionIdx)?.name + " ";
      if (this.#suggestBlob == null || !chars.trim() || this.#suggestBlob?.suggestions.length == 0) {
        return false;
      }
      this.#term.write(removals + chars);
    }
    return "handled";
  }
}
