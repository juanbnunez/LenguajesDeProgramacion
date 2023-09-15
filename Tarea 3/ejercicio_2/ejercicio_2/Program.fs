// Juan Bautista Nunez Exercise 2
let filterStrings (substring: string) (stringList: string list) =
    stringList |> List.filter (fun str -> str.Contains(substring))

let targetSubstring = "Juan"
let stringList = ["Juan juanito juan"; "maria la del barrio"; "Tres tristes tigres comieron trigo en un trigal"; "Juanabanana"]

let result = filterStrings targetSubstring stringList

printfn "%A" result