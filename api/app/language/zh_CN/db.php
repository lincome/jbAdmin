<?php

declare(strict_types=1);

/**
 *  接口返回码各位置说明
 *      1位：应用标识。 0公共标识
 *      2-3位：功能模块标识。   00公共模块；01权限模块
 *      4-6位：错误码。
 */
return [
    'authMenu' => [
        'menuName' => [
            '主页' => '主页',

            '权限管理' => '权限管理',
            '管理员' => '管理员',
            '权限菜单' => '权限菜单',
            '权限组' => '权限组',

            '系统管理' => '系统管理',
            '系统设置' => '系统设置',

            '日志管理' => '日志管理',
            '请求日志' => '请求日志',
        ]
    ]
];
