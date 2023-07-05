function getNumberFromString(str: string) : number {
  const regex = /\d+/g
  const num = str.match(regex)?.join("") || "0"
  return parseInt(num)
}

function loadFullPath(str: string): string {
  if (str.startsWith("/")) return str
  return "/"+str
}

function chopString(s:string, prefix:string, surfix?:string): string {
  const s1 = s.substring(prefix.length)
  if (typeof surfix === 'undefined') return s1
  const s2 = s1.substring(0, s1.length-surfix.length)
  return s2
}

export { getNumberFromString, loadFullPath, chopString }