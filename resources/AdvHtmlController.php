<?php

namespace {{.NameSpace}};

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use {{.ServiceNS}}\{{.ClassName}}Service;
use {{.RepositoryNS}}\{{.ClassName}}Mapper;

#[Route('{{.Route}}')]
class {{.ClassName}}Controller extends AbstractController 
{
    public function __construct(
        private {{.ClassName}}Repository $repository,
        private {{.ClassName}}Service $service,
    )

    #[Route('', methods:['GET'], name: '{{.RouteName}}')]
    public function index(): Response
    {
        return $this->render('{{.TemplateName}}', [
        ]);
    }
}

