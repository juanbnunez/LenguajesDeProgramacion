%Defina sumlist(L, S) que es verdadero si S es la suma de los elementos de L.

sumlist([], 0).

sumlist([H | T], S) :-
    sumlist(T, Ts),
    S is H + Ts.
