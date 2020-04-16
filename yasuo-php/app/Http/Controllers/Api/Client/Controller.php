<?php

namespace App\Http\Controllers\Api\Client;

class Controller extends \App\Http\Controllers\Controller
{
    /**
     * @return mixed
     */
    public function user()
    {
        return $this->request()->user();
    }
}
