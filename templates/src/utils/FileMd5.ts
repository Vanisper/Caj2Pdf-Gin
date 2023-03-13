import SparkMD5 from "spark-md5";

export const FileMd5 = (Uint8Array: Uint8Array, chunkSize?: number) => {
  let startTime = Date.now();
  let usedTime = 0;
  const md5 = new SparkMD5.ArrayBuffer();
  const byteLength = Uint8Array.byteLength;
  let md5Str = "";
  if ((chunkSize && chunkSize > byteLength) || chunkSize === undefined) {
    // 直接处理
    md5.append(Uint8Array);
    md5Str = md5.end();
    usedTime += Date.now() - startTime;
  } else if (chunkSize && chunkSize <= byteLength) {
    let count = parseInt(byteLength / chunkSize + "");
    if (byteLength % chunkSize !== 0) {
      count++;
    }
    for (let index = 0; index < count; index++) {
      const start = index * chunkSize;
      const end = start + chunkSize;
      md5.append(Uint8Array.slice(start, end));
    }
    md5Str = md5.end();
    usedTime += Date.now() - startTime;
  }
  console.log(md5Str, usedTime);
  return md5Str;
};
