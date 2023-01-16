<?php

declare(strict_types=1);

namespace App\Crontab;

class LogRequest
{
    /**
     * 请求日志表分区
     *
     * @return void
     */
    public function partition()
    {
        try {
            dbTablePartition(\App\Module\Db\Dao\Log\Request::class, 24 * 60 * 60, 7);
        } catch (\Throwable $th) {
            //出错时，做记录通知后台管理人员，让其联系技术人工处理。或者不捕获错误，启用app/Listener/CrontabListener监听器处理
            var_dump($th->getMessage());
        }
    }
}
