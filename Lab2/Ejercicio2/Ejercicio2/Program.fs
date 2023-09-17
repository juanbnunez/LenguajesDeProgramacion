//Juan Bautista Nunez Part2
let rec isIsogram (word: string) =
    // Helper function to check if a character appears more than once in the word
    let rec containsDuplicates (charList: char list) =
        match charList with
        | [] -> false
        | x::xs -> List.exists (fun c -> c = x) xs || containsDuplicates xs

    // Convert the word to lowercase and turn it into a list of characters
    let lowercaseWord = word.ToLower().ToCharArray() |> List.ofArray

    // Check if it contains duplicates
    let isIsogram = not (containsDuplicates lowercaseWord)

    // Print the word and the result
    printfn $"Word: %s{word}"
    printfn $"Is an isogram: %b{isIsogram}"

    // Return the result
    isIsogram

let isIsogram2 (word: string) =
    // Helper function to count the occurrences of each letter
    let countOccurrences (charList: char list) =
        charList
        |> List.groupBy (fun c -> c) // Group equal characters
        |> List.map (fun (c, group) -> c, List.length group) // Count the occurrences of each letter

    // Convert the word to lowercase and turn it into a list of characters
    let lowercaseWord = word.ToLower().ToCharArray() |> List.ofArray

    // Count the occurrences of each letter
    let occurrences = countOccurrences lowercaseWord

    // Check if all letters appear exactly twice
    let isIsogram2 = List.forall (fun (_, count) -> count = 2) occurrences

    // Print the word and the result
    printfn $"Word: %s{word}"
    printfn $"Is a second-order isogram: %b{isIsogram2}"

    // Return the result
    isIsogram2

// Example of usage
let result = isIsogram "Juan"
let result2 = isIsogram2 "deed"


