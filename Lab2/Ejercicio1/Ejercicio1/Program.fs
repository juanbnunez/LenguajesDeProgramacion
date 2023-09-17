//Juan Bautista Nunez Part 1
// Function to calculate the percentage of accuracy between two chromosomes
let chromosomeSubjectCandidate (subject: List<int>) (candidate: List<int>) =
    let matches = subject |> List.zip candidate |> List.filter (fun (s, c) -> s = c) |> List.length
    float matches * 100.0 / float subject.Length

// Function to identify the subject most similar to a sample from a list of people
let mostSimilarSubject (sample: List<int>) (people: List<List<int>>) =
    let similarSubject =
        people
        |> List.maxBy (fun person -> chromosomeSubjectCandidate sample person)
    let similarityPercentage =
        chromosomeSubjectCandidate sample similarSubject
    (similarSubject, similarityPercentage)

// Example of using the chromosomeSubjectCandidate function
let subject = [1; 2; 3; 4; 5; 6; 7; 8; 9; 10]
let candidate = [77; 2; 3; 4; 5; 7; 8; 33; 10; 50]
let matches = chromosomeSubjectCandidate subject candidate
printfn $"Percentage of match between the sample and the candidate: %f{matches}"

// Example of using the mostSimilarSubject function
let sample = [1; 2; 3; 4; 5; 6; 7; 8; 9; 10]
let people = [
    [1; 2; 3; 4; 5; 6; 7; 8; 88; 10];
    [1; 2; 3; 4; 5; 7; 8; 9; 10; 11];
    [1; 2; 3; 4; 5; 6; 7; 8; 9; 11];
    [1; 2; 3; 4; 5; 6; 7; 8; 9; 12]
]
let similarSubject, similarityPercentage = mostSimilarSubject sample people
printfn $"Sample: %A{sample}"
printfn $"Most similar person: %A{similarSubject}"
printfn $"Similarity percentage: %f{similarityPercentage}"
