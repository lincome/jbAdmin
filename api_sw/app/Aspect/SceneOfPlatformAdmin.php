<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Aop\ProceedingJoinPoint;

//#[\Hyperf\Di\Annotation\Aspect]
class SceneOfPlatformAdmin extends \Hyperf\Di\Aop\AbstractAspect
{
    #[\Hyperf\Di\Annotation\Inject]
    protected \Psr\Container\ContainerInterface $container;

    //执行优先级（大值优先）
    public ?int $priority = 19;

    //切入的类
    public array $classes = [
        \App\Controller\Login::class . '::info',
        \App\Controller\Login::class . '::updateInfo',
        \App\Controller\Login::class . '::menuTree',
        \App\Controller\Auth\Action::class,
        \App\Controller\Auth\Menu::class,
        \App\Controller\Auth\Role::class,
        \App\Controller\Auth\Scene::class,
        \App\Controller\Log\Request::class,
        \App\Controller\Platform\Admin::class,
        \App\Controller\Platform\Config::class,
        \App\Controller\Platform\Server::class,
    ];

    //切入的注解
    public array $annotations = [];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        try {
            $sceneCode = $this->container->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneCode();
            if ($sceneCode == 'platformAdmin') {
                $this->container->get(\App\Module\Service\Login::class)->verifyToken($sceneCode);
            }
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
