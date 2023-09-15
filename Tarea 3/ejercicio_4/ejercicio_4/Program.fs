// Juan Bautista Nunez Exercise 4
type Node = (string * string list * int list)
let private graph1 = [
    ("A", ["B"; "C"], [10; 20]);
    ("B", ["A"; "C"; "D"], [5; 15; 25]);
    ("C", ["A"; "B"; "D"], [12; 8; 30]);
    ("D", ["B"; "C"], [3; 9]);
]
let private graph2 = [
    ("X", ["Y"; "Z"], [7; 14]);
    ("Y", ["X"; "Z"], [5; 12]);
    ("Z", ["X"; "Y"], [6; 11]);
]
type Graph = Node list

let rec neighbors (node: string) = function
    | [] -> []
    | (head, neigh, _) :: rest ->
        if head = node then neigh
        else neighbors node rest

let membr elem list =
    List.exists (fun x -> x = elem) list

let extend path graph = 
    List.filter
        (fun x -> x <> [])
        (List.map (fun x -> if (membr x path) then [] else x::path) (neighbors (List.head path) graph)) 

let rec getWeight from to_ = function
    | (start, path, weight)::rest when start = from ->
        List.tryFindIndex (fun x -> x = to_) path
        |> Option.bind (fun index -> if List.item index weight <> 0 then Some (List.item index weight) else None)
        |> Option.defaultValue 0
    | _::rest -> getWeight from to_ rest
    | [] -> 0

let calculateWeight (graph: Graph) =
    fun path ->
        let rec accumulateWeight accum = function
            | [] | [_] -> accum
            | x::y::rest -> accumulateWeight (accum + getWeight x y graph) (y::rest)
        (path, accumulateWeight 0 path)

let shortestPath start finish graph =
    let rec depthAux finish path graph =
        if List.isEmpty path then []
        elif (List.head (List.head path)) = finish then
            List.rev (List.head path) :: depthAux finish ((extend (List.head path) graph) @ (List.tail path)) graph       
        else
            depthAux finish ((List.tail path) @ (extend (List.head path) graph)) graph
    depthAux finish [[start]] graph |> List.map (calculateWeight graph) |> List.minBy (fun n -> snd n) |> fst

[<EntryPoint>]
let main (args: string array): int =
    printf "Shortest path in graph 1: \n"
    shortestPath "A" "D" graph1 |> printfn "%A"
    printf "Shortest path in graph 2: \n"
    shortestPath "X" "Z" graph2 |> printfn "%A"
    0
