<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Logic\Auth\Menu as LogicAuthMenu;
use App\Module\Service\AbstractService;

class Menu extends AbstractService
{
    /**
     * 创建
     *
     * @param array $data
     * @return void
     */
    public function create(array $data)
    {
        if (!empty($data['pid'])) {
            $pInfo = $this->getDao()->field(['pidPath', 'level'])->where(['id' => $data['pid'], 'sceneId' => $data['sceneId']])->getInfo();
            if (empty($pInfo)) {
                throwFailJson('999303');
            }
        }
        $id = $this->getDao()->insert($data)->saveInsert();
        if (empty($id)) {
            throwFailJson('999999');
        }
        if (!empty($data['pid'])) {
            $update['pidPath'] = $pInfo->pidPath . '-' . $id;
            $update['level'] = $pInfo->level + 1;
        } else {
            $update['pidPath'] = '0-' . $id;
            $update['level'] = 1;
        }
        $this->getDao()->update($update)->where(['id' => $id])->saveUpdate();
        throwSuccessJson();
    }

    /**
     * 更新
     *
     * @param array $data
     * @param array $where
     * @return void
     */
    public function update(array $data, array $where)
    {
        if (isset($data['pid'])) {
            $oldInfo = $this->getDao()->where($where)->getInfo();
            if ($data['pid'] == $oldInfo->menuId) {
                throwFailJson('999304');
            }
            if ($data['pid'] == $oldInfo->pid) {
                unset($data['pid']);
            } else {
                if ($data['pid'] > 0) {
                    $pInfo = $this->getDao()->field(['pidPath', 'level'])->where(['id' => $data['pid'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getInfo();
                    if (empty($pInfo)) {
                        throwFailJson('999303');
                    }
                    if (in_array($oldInfo->menuId, explode('-',  $pInfo->pidPath))) {
                        throwFailJson('999305');
                    }
                    $data['pidPath'] =  $pInfo->pidPath . '-' . $oldInfo->menuId;
                    $data['level'] = $pInfo->level + 1;
                } else {
                    $data['pidPath'] = '0-' . $oldInfo->menuId;
                    $data['level'] = 1;
                }
            }
        }
        $result = $this->getDao()->where($where)->update($data)->saveUpdate();
        if (empty($result)) {
            throwFailJson('999999');
        }
        //修改pid时，更新所有子孙级的pidPath和level
        if (isset($data['pid'])) {
            $this->getDao()->where([['pidPath', 'like', $oldInfo->pidPath . '%']])
                ->update([
                    'pidPathOfChild' => [$data['pidPath'], $oldInfo->pidPath],
                    'levelOfChild' => $data['level'] - $oldInfo->level,
                ])
                ->saveUpdate();
        }
        throwSuccessJson();
    }

    /**
     * 删除
     *
     * @param array $where
     * @return void
     */
    public function delete(array $where)
    {
        $idArr = $where['id'] ?? $this->getDao()->where($where)->getBuilder()->pluck('menuId')->toArray();
        if ($this->getDao()->where(['pid' => $idArr])->getBuilder()->exists()) {
            throwFailJson('999306');
        }
        $result = $this->getDao()->where($where)->delete();
        if (empty($result)) {
            throwFailJson('999999');
        }
        throwSuccessJson();
    }

    /**
     * 获取树状权限菜单
     *
     * @param array $field
     * @param array $where
     * @return void
     */
    public function tree(array $field = [], array $where = [])
    {
        $list = $this->getDao()->field($field)->where($where)->getList();

        $tree = $this->container->get(LogicAuthMenu::class)->tree($list);
        throwSuccessJson(['tree' => $tree]);
    }
}
