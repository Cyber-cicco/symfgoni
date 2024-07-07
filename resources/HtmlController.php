<?php

namespace {{.NameSpace}};

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

#[Route('{{.Route}}')]
class {{.ClassName}}Controller extends AbstractController 
{
    #[Route('', methods:['GET'], name: '{{.RouteName}}')]
    public function index(): Response
    {
        return $this->render('{{.TemplateName}}', [
        ]);
    }
}

