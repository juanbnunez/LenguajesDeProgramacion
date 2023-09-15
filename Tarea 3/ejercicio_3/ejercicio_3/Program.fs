// Juan Bautista Nunez Exercise 3

let nthElement n list =
    let indices = [0..(List.length list - 1)]
    let elements = List.map2 (fun i x -> (i, x)) indices list
    match List.tryFind (fun (i, _) -> i = n) elements with
    | Some (_, element) -> Some element
    | None -> None

let result1 = nthElement 2 [1; 2; 3; 4; 5]
let result2 = nthElement 3 [1; 2; 3; 4; 5]
let result3 = nthElement 6 [1; 2; 3; 4; 5]

match result1 with
| Some value -> printfn "%A" value
| None -> printfn "False"

match result2 with
| Some value -> printfn "%A" value
| None -> printfn "False"

match result3 with
| Some value -> printfn "%A" value
| None -> printfn "False"