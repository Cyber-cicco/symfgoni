<?php

namespace {{.NameSpace}};

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Attribute\Route;
use {{.ServiceNS}}\{{.ClassName}}Service;
use {{.MapperNS}}\{{.ClassName}}Mapper;

#[Route('{{.Route}}')]
class {{.ClassName}}Controller extends AbstractController 
{
    public function __construct(
        private {{.ClassName}}Service $service,
        private {{.ClassName}}Mapper $repo,
    ){
    }
    
    #[Route('', methods:['GET'], name: '{{.RouteName}}')]
    public function get(Request $req): array
    {
        return ["Weeeeeeeeeeeee"];
    }
}
