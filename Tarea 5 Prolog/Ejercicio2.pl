
person("Person1", "Lastname1", [1, 0, 0, 0, 1, 0, 0, 0, 1, 1]).
person("Person2", "Lastname2", [1, 1, 1, 1, 0, 0, 0, 0, 0, 0]).
person("Person3", "Lastname3", [0, 0, 0, 0, 1, 1, 1, 1, 1, 1]).
person("Person4", "Lastname4", [1, 0, 1, 0, 1, 0, 1, 0, 1, 0]).
person("Person5", "Lastname5", [0, 1, 0, 1, 0, 1, 0, 1, 0, 1]).
person("Person6", "Lastname6", [1, 1, 0, 0, 1, 1, 0, 0, 1, 1]).
person("Person7", "Lastname7", [0, 0, 1, 1, 0, 0, 1, 1, 0, 0]).
person("Person8", "Lastname8", [1, 1, 1, 0, 0, 0, 1, 1, 1, 0]).
person("Person9", "Lastname9", [0, 0, 0, 1, 1, 1, 0, 0, 0, 1]).
person("Person10", "Lastname10", [1, 0, 0, 1, 1, 0, 0, 1, 1, 0]).

% calculate accuracy between two chromosomes
accuracy_percentage(SubjectChromosome, CandidateChromosome, Percentage) :-
    length(SubjectChromosome, Length),
    accuracy(SubjectChromosome, CandidateChromosome, 0, Length, Percentage).

accuracy(_, _, Accumulator, 0, Accumulator).  % Base case
accuracy([X|RestSubject], [X|RestCandidate], Accumulator, N, Accuracy) :-
    N > 0,
    NewAccumulator is Accumulator + 1,
    N1 is N - 1,
    accuracy(RestSubject, RestCandidate, NewAccumulator, N1, Accuracy).
accuracy([_|RestSubject], [_|RestCandidate], Accumulator, N, Accuracy) :-
    N > 0,
    N1 is N - 1,
    accuracy(RestSubject, RestCandidate, Accumulator, N1, Accuracy).

% most similar to a sample
most_similar_subject(Sample, People, SimilarSubject, SimilarityPercentage) :-
    similar_subject_aux(Sample, People, _, -1, SimilarSubject, SimilarityPercentage).

similar_subject_aux(_, [], SimilarSubject, SimilarityPercentage, SimilarSubject, SimilarityPercentage).  % Base case
similar_subject_aux(Sample, [Person|RestPeople], _, CurrentPercentage, SimilarSubject, SimilarityPercentage) :-
    accuracy_percentage(Sample, Person, Percentage),
    Percentage >= CurrentPercentage,
    similar_subject_aux(Sample, RestPeople, Person, Percentage, SimilarSubject, SimilarityPercentage).
similar_subject_aux(Sample, [Person|RestPeople], CurrentSubject, CurrentPercentage, SimilarSubject, SimilarityPercentage) :-
    accuracy_percentage(Sample, Person, Percentage),
    Percentage < CurrentPercentage,
    similar_subject_aux(Sample, RestPeople, CurrentSubject, CurrentPercentage, SimilarSubject, SimilarityPercentage).

% Sample
sample([0, 0, 0, 1, 1, 1, 0, 0, 0, 1]).

% most similar to the sample
find_similar_subject(SimilarSubject, SimilarityPercentage) :-
    sample(Sample),
    findall((Name, Lastname, Percentage), (person(Name, Lastname, Chromosome), accuracy_percentage(Sample, Chromosome, Percentage)), SimilarityList),
    max_similarity(SimilarityList, SimilarSubject, SimilarityPercentage).

%maximum similarity 
max_similarity([(Name, _, Percentage)], Name, Percentage).
max_similarity([(Name, _, Percentage)|Rest], SimilarSubject, SimilarityPercentage) :-
    max_similarity(Rest, OtherName, OtherPercentage),
    (Percentage >= OtherPercentage ->
        SimilarSubject = Name,
        SimilarityPercentage = Percentage
    ;
        SimilarSubject = OtherName,
        SimilarityPercentage = OtherPercentage
    ).

% Use exaple
%?- find_similar_subject(SimilarSubject, SimilarityPercentage).
