<?php

declare(strict_types=1);

namespace App\Controller\Platform;

use App\Controller\AbstractController;

class Config extends AbstractController
{
    /**
     * 获取
     *
     * @return void
     */
    public function get()
    {
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__);
                $this->checkAuth(__FUNCTION__, $sceneCode);
                $this->service->get(['configKey' => $data['configKeyArr']]);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }

    /**
     * 保存
     *
     * @return void
     */
    public function save()
    {
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__);
                $this->checkAuth(__FUNCTION__, $sceneCode);
                $this->service->save($data);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }
}
