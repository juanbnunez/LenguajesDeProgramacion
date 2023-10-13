connected(i, a, 2).
connected(i, b, 5).
connected(a, b, 3).
connected(a, c, 1).
connected(b, f, 4).
connected(b, c, 2).
connected(c, f, 3).

connected(X, Y, Weight) :- connected(X, Y, Weight).
connected(X, Y, Weight) :- connected(Y, X, Weight).

path(Start, End, Route, Weight) :- path2(Start, End, [], Route, 0, Weight).

path2(End, End, _, [End], AccumulatedWeight, TotalWeight) :-
    TotalWeight is AccumulatedWeight.
path2(Start, End, Visited, [Start | Rest], AccumulatedWeight, TotalWeight) :-
    connected(Start, Someone, Weight),
    not(member(Someone, Visited)),
    NewAccumulatedWeight is AccumulatedWeight + Weight,
    path2(Someone, End, [Start | Visited], Rest, NewAccumulatedWeight, TotalWeight).

%Use exaple
%?- path(i, f, Route, Weight).
