sub_cadenas(_, [], []).

sub_cadenas(Subcadena, [Cadena|Resto], Filtradas) :-
    sub_atom(Cadena, _, _, _, Subcadena),
    sub_cadenas(Subcadena, Resto, FiltradasRestantes),
    Filtradas = [Cadena|FiltradasRestantes].

sub_cadenas(Subcadena, [_|Resto], Filtradas) :-
    sub_cadenas(Subcadena, Resto, Filtradas).
