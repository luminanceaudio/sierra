import path from 'path';
import { spawn } from 'child_process';
import { ChildProcess } from 'node:child_process';
import { getBinariesPath } from '../util';

let serverProcess: ChildProcess;

function spawnServer(args: string[]) {
  const sierraBinaryFilePath = path.join(getBinariesPath(), 'sierraserver');

  console.log('Starting server ', sierraBinaryFilePath);
  serverProcess = spawn(sierraBinaryFilePath, args, {
    stdio: 'inherit',
  });
}

export function startServer() {
  spawnServer(['server', 'start']);

  serverProcess.on('exit', (): void => {
    spawnServer(['server', 'start']);
  });
}

export function terminateServer() {
  spawnServer(['server', 'shutdown']);
}
