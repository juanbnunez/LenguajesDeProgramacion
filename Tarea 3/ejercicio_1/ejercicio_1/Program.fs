// Juan Bautista Nunez Exercise 1

let rec shiftElementsLeft n list =
    match n with
    | 0 -> list
    | _ ->
        match list with
        | [] -> []
        | hd::tl -> shiftElementsLeft (n - 1) (tl @ [0])

let rec shiftElementsRight n list =
    match n with
    | 0 -> list
    | _ ->
        match list with
        | [] -> []
        | lst ->
            let rest = List.init (List.length lst - 1) (fun i -> List.item i lst)
            shiftElementsRight (n - 1) (0::rest)

let shiftElements direction n list =
    match direction with
    | "left" ->
         shiftElementsLeft n list
    | "right" ->
         shiftElementsRight n list
    | _ -> failwith "Invalid direction"

let originalList = [1; 2; 3; 4; 5]

let leftResult = shiftElements "left" 3 originalList
printfn "Shift left by 3 positions: %A" leftResult

let rightResult = shiftElements "right" 2 originalList
printfn "Shift right by 2 positions: %A" rightResult

let leftResult2 = shiftElements "left" 6 originalList
printfn "Shift left by 6 positions: %A" leftResult2