<?php

namespace {{.NameSpace}};

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Attribute\Route;

#[Route('{{.Route}}')]
class {{.ClassName}}Controller extends AbstractController 
{
    #[Route('', methods:['GET'], name: '{{.RouteName}}')]
    public function get(Request $req): array
    {
        return ["Weeeeeeeeeeeee"];
    }
}
