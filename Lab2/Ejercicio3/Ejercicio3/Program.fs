//Juan Bautista Nunez Part3

// Define the matrix
let matrix: char[][] = 
    [|
        [| 'A'; 'B'; 'C'; 'D'; 'E'; 'F' |];
        [| 'G'; 'J'; 'U'; 'A'; 'N'; 'L' |];
        [| 'L'; 'M'; 'N'; 'S'; 'T'; 'U' |];
        [| 'V'; 'W'; 'X'; 'Y'; 'Z'; 'A' |];
        [| 'A'; 'S'; 'A'; 'D'; 'O'; 'G' |];
        [| 'D'; 'E'; 'P'; 'E'; 'R'; 'O' |]
    |]

// Get the number of columns and rows in the matrix.
let columns = Array.length matrix[0]
let rows = Array.length matrix

// Function to get a letter at a given coordinate (x, y) in the matrix.
let getLetter (x, y) =
    matrix[x].[y]

// Function to check if an element exists in a list.
let existsInList (elem: 'a) (list: 'a list) =
    List.exists (fun x -> x = elem) list

// Function to print the coordinates of a word in the matrix.
let printCoordinates (word: string) (coordinates: (int * int) list) =
    printfn $"Coordinates of the word \"%s{word}\":"
    for x, y in coordinates do
        printfn $"Letter: %c{getLetter (x, y)} - Coordinates: (%d{x}, %d{y})"

// Recursive function to search for words starting from a given coordinate.
let rec searchWords (word: string) (coord: int * int) (visited: (int * int) list) =
    let x, y = coord
    if x >= 0 && x < rows && y >= 0 && y < columns && not (existsInList (x, y) visited) then
        let currentLetter = getLetter (x, y)
        if currentLetter = word[0] then
            if word.Length = 1 then
                [[(x, y)]]
            else
                let visited' = (x, y) :: visited
                let neighbors = [(x + 1, y); (x - 1, y); (x, y + 1); (x, y - 1);
                                 (x + 1, y + 1); (x + 1, y - 1); (x - 1, y + 1); (x - 1, y - 1)]
                let results =
                    neighbors
                    |> List.collect (fun (nx, ny) -> searchWords (word.Substring(1)) (nx, ny) visited')
                if List.length results > 0 then
                    List.map (fun list -> (x, y) :: list) results
                else
                    []
        else
            []
    else
        []

// Function to find a word in the matrix.
let findWord (word: string) =
    let word = word.ToUpper()
    let results =
        [ for x in 0 .. rows - 1 do
              for y in 0 .. columns - 1 do
                  let result = searchWords word (x, y) []
                  if List.length result > 0 then
                      yield result
        ]
    let finalResults =
        results
        |> List.collect (fun list -> list)
    if List.length finalResults > 0 then
        finalResults
    else
        []

// Define the words to find.
let wordToFind1 = "JUAN"
let wordToFind2 = "SUEÑO"
let wordToFind3 = "ASADO"

// Search for the words.
let result1 = findWord wordToFind1
let result2 = findWord wordToFind2
let result3 = findWord wordToFind3

// Print the coordinates of each word if found, otherwise, print a message.
if List.length result1 > 0 then
    printCoordinates wordToFind1 (List.head result1)
else
    printfn $"The word \"%s{wordToFind1}\" was not found in the word search puzzle."

if List.length result2 > 0 then
    printCoordinates wordToFind2 (List.head result2)
else
    printfn $"The word \"%s{wordToFind2}\" was not found in the word search puzzle."

if List.length result3 > 0 then
    printCoordinates wordToFind3 (List.head result3)
else
    printfn $"The word \"%s{wordToFind3}\" was not found in the word search puzzle."
