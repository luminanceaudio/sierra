import path from 'path';
import { spawn } from 'child_process';
import { ChildProcessWithoutNullStreams } from 'node:child_process';
import { getBinariesPath } from '../util';

let serverProcess: ChildProcessWithoutNullStreams;

export function startServer() {
  const sierraBinaryFilePath = path.join(getBinariesPath(), 'sierra');

  console.log('Starting server ', sierraBinaryFilePath);
  serverProcess = spawn(sierraBinaryFilePath, ['server']);

  serverProcess.on('exit', (code: number): void => {
    // Re-run
    serverProcess = spawn(`${sierraBinaryFilePath} server`);
  });

  serverProcess.on('spawn', (code: number): void => {
    console.log(`Spawn: ${code}`);
  });

  serverProcess.on('error', (err: Error): void => {
    console.error(err);
  });
}
