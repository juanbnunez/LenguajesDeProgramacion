// Juan Bautista Nunez Exercise 5
let graph1 = [
    ("Start", ["2"]);
    ("2", ["3"; "8"]);
    ("8", ["2"; "9"]);
    ("3", ["2"; "9"; "4"]);
    ("9", ["8"; "3"]);
    ("4", ["3"; "10"]);
    ("10", ["4"; "16"]);
    ("16", ["10"; "22"]);
    ("22", ["16"; "21"]);
    ("21", ["22"; "15"]);
    ("15", ["21"; "14"]);
    ("14", ["15"; "20"; "13"]);
    ("13", ["14"; "7"; "1"]);
    ("7", ["13"; "1"]);
    ("1", ["7";]);
    ("20", ["14"; "26"]);
    ("26", ["20"; "27"]);
    ("27", ["26"; "28"]);
    ("28", ["27"; "34"; "29"]);
    ("29", ["28"; "23"]);
    ("23", ["29"; "17"]);
    ("17", ["23"; "11"]);
    ("11", ["17"; "5"]);
    ("5", ["11"; "6"]);
    ("6", ["5"]);
    ("34", ["28"; "33"; "35"]);
    ("33", ["34"; "32"]);
    ("32", ["33"; "31"; "End"]);
    ("31", ["32"; "25"]);
    ("25", ["31"; "19"]);
    ("19", ["25"]);
    ("End", ["32"]);
    ("35", ["34"; "36"]);
    ("36", ["35"; "30"]);
    ("30", ["36"; "24"]);
    ("24", ["30"; "18"]);
    ("18", ["24"; "12"]);
    ("12", ["18"])
]

//Without wall 8 to 14
let graph2 = [
    ("Start", ["2"]);
    ("2", ["3"; "8"]);
    ("8", ["2"; "9"; "14"]);
    ("3", ["2"; "9"; "4"]);
    ("9", ["8"; "3"]);
    ("4", ["3"; "10"]);
    ("10", ["4"; "16"]);
    ("16", ["10"; "22"]);
    ("22", ["16"; "21"]);
    ("21", ["22"; "15"]);
    ("15", ["21"; "14"]);
    ("14", ["15"; "20"; "13"; "8"]);
    ("13", ["14"; "7"; "1"]);
    ("7", ["13"; "1"]);
    ("1", ["7";]);
    ("20", ["14"; "26"]);
    ("26", ["20"; "27"]);
    ("27", ["26"; "28"]);
    ("28", ["27"; "34"; "29"]);
    ("29", ["28"; "23"]);
    ("23", ["29"; "17"]);
    ("17", ["23"; "11"]);
    ("11", ["17"; "5"]);
    ("5", ["11"; "6"]);
    ("6", ["5"]);
    ("34", ["28"; "33"; "35"]);
    ("33", ["34"; "32"]);
    ("32", ["33"; "31"; "End"]);
    ("31", ["32"; "25"]);
    ("25", ["31"; "19"]);
    ("19", ["25"]);
    ("End", ["32"]);
    ("35", ["34"; "36"]);
    ("36", ["35"; "30"]);
    ("30", ["36"; "24"]);
    ("24", ["30"; "18"]);
    ("18", ["24"; "12"]);
    ("12", ["18"])
]
let miembro elem lista =
    List.exists (fun x -> x = elem) lista

let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (head, neighbors) :: rest ->
        if head = nodo then neighbors
        else vecinos nodo rest

let extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map  (fun x -> if (miembro x ruta) then [] else x::ruta) (vecinos (List.head ruta) grafo)) 

let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) :: prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)) grafo
    let resultado = prof_aux fin [[ini]] grafo
    List.iter (fun ruta -> printfn "Ruta: %A" ruta) resultado
    resultado

let shortestPath ini fin grafo =
    let rec bfs queue visited =
        match queue with
        | [] -> None
        | (currentNode, path) :: rest ->
            if currentNode = fin then
                Some (List.rev (currentNode :: path)) 
            else
                let neighbors = vecinos currentNode grafo
                let newPaths =
                    neighbors
                    |> List.filter (fun neighbor -> not (miembro neighbor visited))
                    |> List.map (fun neighbor -> (neighbor, currentNode :: path))
                bfs (rest @ newPaths) (currentNode :: visited)
    
    let initialQueue = [(ini, [])] 
    match bfs initialQueue [] with
    | Some resultPath -> resultPath
    | None -> [] 



printfn"Original route: \n"
let ruta1 = prof2 "Start" "End" graph1
printfn"\n"
printfn"Ruta quitando la pared del 22 hacia el 28: \n"
let ruta2 = prof2 "Start" "End" graph2

printfn"\n"
printfn "Shortest route to origin: \n"
let ruta = shortestPath "Start" "End" graph1
match ruta with
| [] -> printfn "Cant find route."
| _ -> printfn "Route finded: %A" ruta
