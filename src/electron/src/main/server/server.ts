import path from 'path';
import { spawn } from 'child_process';
import { ChildProcessWithoutNullStreams } from 'node:child_process';
import { getBinariesPath } from '../util';

let serverProcess: ChildProcessWithoutNullStreams;

export function startServer() {
  const sierraBinaryFilePath = path.join(getBinariesPath(), 'sierraserver');

  console.log('Starting server ', sierraBinaryFilePath);
  serverProcess = spawn(sierraBinaryFilePath, ['server', 'start']);

  serverProcess.on('exit', (): void => {
    // Re-run
    serverProcess = spawn(sierraBinaryFilePath, ['server', 'start']);
  });

  serverProcess.on('error', (err: Error): void => {
    console.error(err);
  });
}
