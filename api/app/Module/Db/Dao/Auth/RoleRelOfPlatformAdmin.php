<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $roleId 权限角色ID
 * @property int $adminId 平台管理员ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class RoleRelOfPlatformAdmin extends AbstractDao
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_role_rel_of_platform_admin';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'default';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['roleId', 'adminId', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['roleId' => 'integer', 'adminId' => 'integer'];
}
