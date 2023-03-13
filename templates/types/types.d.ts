declare interface Window {
  getFileMd5: (Uint8Array: Uint8Array, chunkSize?: number) => string;
  wasmMd5Add: (Uint8Array: Uint8Array) => void;
  wasmMd5End: () => string;
}
