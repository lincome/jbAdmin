<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $roleId 权限角色ID
 * @property int $adminId 平台管理员ID
 * @property string $updateAt 更新时间
 * @property string $createAt 创建时间
 */
class RoleRelOfPlatformAdmin extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_role_rel_of_platform_admin';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['roleId', 'adminId', 'updateAt', 'createAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['roleId' => 'integer', 'adminId' => 'integer'];
}
