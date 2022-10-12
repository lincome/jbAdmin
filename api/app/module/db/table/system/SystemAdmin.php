<?php

declare(strict_types=1);

namespace app\module\db\table\system;

use app\module\db\table\AbstractTable;
use DI\Annotation\Inject;
use app\module\db\table\auth\AuthRole;
use app\module\db\table\auth\AuthRoleRelOfSystemAdmin;

class SystemAdmin extends AbstractTable
{
    /**
     * @Inject
     * @var \app\module\db\model\system\SystemAdmin
     */
    protected $model;

    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return self
     */
    protected function fieldOfAlone(string $key): self
    {
        switch ($key) {
            case 'roleName':
                $this->joinOfAlone($key);
                $this->field['select'][] = container(AuthRole::class, true)->getTableAlias() . '.' . $key;
                break;
            default:
                $this->field['select'][] = $key;
                break;
        }
        return $this;
    }

    /**
     * 解析join（独有的）
     *
     * @param string $key   键，用于确定关联表
     * @param [type] $value 值，用于确定关联表
     * @return self
     */
    protected function joinOfAlone(string $key, $value = null): self
    {
        switch ($key) {
            case 'roleName':
                $tableAuthRoleRelOfSystemAdmin = container(AuthRoleRelOfSystemAdmin::class, true);
                $tableAuthRoleRelOfSystemAdminAlias = $tableAuthRoleRelOfSystemAdmin->getTableAlias();
                if (!isset($this->join[$tableAuthRoleRelOfSystemAdminAlias])) {
                    $this->join[$tableAuthRoleRelOfSystemAdminAlias] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $tableAuthRoleRelOfSystemAdmin->getTable() . ' AS ' . $tableAuthRoleRelOfSystemAdminAlias,
                            $tableAuthRoleRelOfSystemAdminAlias . '.adminId',
                            '=',
                            $this->getTableAlias() . '.' . $this->getPrimaryKey()
                        ]
                    ];
                }
                $tableAuthRole = container(AuthRole::class, true);
                $tableAuthRoleAlias = $tableAuthRole->getTableAlias();
                if (!isset($this->join[$tableAuthRoleAlias])) {
                    $this->join[$tableAuthRoleAlias] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $tableAuthRole->getTable() . ' AS ' . $tableAuthRoleAlias,
                            $tableAuthRoleAlias . '.roleId',
                            '=',
                            $tableAuthRoleRelOfSystemAdminAlias . '.roleId'
                        ]
                    ];
                }
                break;
        }
        return $this;
    }
}
