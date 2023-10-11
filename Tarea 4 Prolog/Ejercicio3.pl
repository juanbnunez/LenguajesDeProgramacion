aplanar([],[]).

% Caso 1: Si el primer elemento de la lista es una lista
aplanar([X|Resto], L2) :-
    is_list(X),  
    aplanar(X, XAplanada), 
    aplanar(Resto, RestoAplanado),  
    append(XAplanada, RestoAplanado, L2).

% Caso 2: Si el primer elemento de la lista no es una lista
aplanar([X|Resto], [X|RestoAplanado]) :-
    \+ is_list(X), 
    aplanar(Resto, RestoAplanado).