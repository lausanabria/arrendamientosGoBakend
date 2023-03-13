create DATABASE arriendos;

use arriendos;

create TABLE pagos(
    documentoIdentificacionArrendatario int not null,
    codigoInmueble VARCHAR(50) not null,
    valorPagado int,
    fechaPago date,
    primary key(documentoIdentificacionArrendatario,codigoInmueble)
)ENGINE = INNODB;
