/* eslint import/prefer-default-export: off */
import { URL } from 'url';
import path from 'path';
import { platform, arch } from 'os';
import { app } from 'electron';

export function resolveHtmlPath(htmlFileName: string) {
  if (process.env.NODE_ENV === 'development') {
    const port = process.env.PORT || 1212;
    const url = new URL(`http://localhost:${port}`);
    url.pathname = htmlFileName;
    return url.href;
  }
  return `file://${path.resolve(__dirname, '../renderer/', htmlFileName)}`;
}

export function getPlatform(): string {
  switch (platform()) {
    case 'aix':
    case 'freebsd':
    case 'linux':
    case 'openbsd':
    case 'android':
      return 'linux';
    case 'darwin':
    case 'sunos':
      return 'mac';
    case 'win32':
      return 'win';
    default:
      console.error(`Unsupported platform: ${platform()}`);
      // Fallback to assume linux
      return 'linux';
  }
}

export function getBinariesPath() {
  let base = process.resourcesPath;
  if (process.env.NODE_ENV === 'development' || !app.isPackaged) {
    base = app.getAppPath();
  }

  return path.join(base, 'binaries', getPlatform()!, arch());
}
