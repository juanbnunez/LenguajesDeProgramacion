partir([], _, [], []).

partir([X|Resto], Umbral, [X|Menores], Mayores) :-
    X =< Umbral, % X es menor o igual al umbral
    partir(Resto, Umbral, Menores, Mayores).

partir([X|Resto], Umbral, Menores, [X|Mayores]) :-
    X > Umbral, % X es mayor al umbral
    partir(Resto, Umbral, Menores, Mayores).
