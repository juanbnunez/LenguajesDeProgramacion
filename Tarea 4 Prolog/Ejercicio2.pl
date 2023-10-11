%Defina la relación subconj(S, S1), donde S y S1 son listas 
%representando conjuntos, que es verdadera si S1 es subconjunto de S.

subconj([], _).

% Caso recursivo: 
subconj([X|RestoS1], S) :-
    member(X, S),
    subconj(RestoS1, S).
