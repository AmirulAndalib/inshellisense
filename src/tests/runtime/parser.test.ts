// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import { parseCommand } from "../../runtime/parser.js";
import { Shell } from "../../utils/shell.js";

const testData = [
  { command: `cmd --flag value` },
  { command: `cmd --flag=value` },
  { command: `cmd --flag='value' ` },
  { command: `cmd --flag="value" ` },
  { command: `cmd 'value' ` },
  { command: `cmd value ` },
  { command: `cmd -f` },
  { command: `cmd -f=value ` },
  { command: `cmd -f value ` },
  { command: `cmd -f 'value' ` },
  { command: `cmd -f="value" ` },
  { command: `cmd -f='val` },
  { command: `cmd -f` },
  { command: `cmd -f=` },
  { command: `cmd -f ` },
  { command: `cmd` },
  { command: `cmd ` },
  { command: `cmd "value' ` },
  { command: `cmd "value" ` },
  { command: `cmd "value'\\"\\"" ` },
  { command: `cmd1 | cmd2 ` },
  { command: `cmd1 -` },
  { command: `cmd dir\\ 1/dir\\ 2/item1` },
  { command: `cmd1 "item1"item2` },
  { command: `cmd1 "item1"item2 item3` },
  { command: `cmd1 "item1"item2 "item3"` },
  { command: "`cmd1`" },
  { command: "cmd 'item1\\item2\\'item3 ", shell: Shell.Powershell },
  { command: "cmd 'item1\\item2\\'item3", shell: Shell.Powershell },
  { command: "😁" },
];

describe(`parseCommand`, () => {
  testData.forEach(({ command, shell }) => {
    test(command, () => {
      expect(parseCommand(command, shell ?? Shell.Bash)).toMatchSnapshot();
    });
  });
});
