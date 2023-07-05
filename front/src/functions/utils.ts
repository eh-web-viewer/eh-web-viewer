function getNumberFromString(str: string) : number {
  const regex = /[\d]+/g
  const num = str.match(regex)?.join("") || "0"
  return parseInt(num)
}

function loadFullPath(str: string): string {
  if (str.startsWith("/")) return str
  return "/"+str
}

export { loadFullPath }