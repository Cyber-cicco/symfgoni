<?php

namespace {{.NameSpace}};

use {{.RepositoryNS}}\{{.ClassName}}Repository;
use {{.MapperNS}}\{{.ClassName}}Mapper;

class {{.ClassName}}Service
{

    public function __construct(
        private {{.ClassName}}Repository $repo,
        private {{.ClassName}}Mapper $mapper,
    ){
    }

}
