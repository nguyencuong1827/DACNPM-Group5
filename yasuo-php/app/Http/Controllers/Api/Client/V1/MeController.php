<?php

namespace App\Http\Controllers\Api\Client\V1;

use App\Http\Controllers\Api\Client\Controller;

class MeController extends Controller
{
    public function logout()
    {
        return $this->user()->logout();
    }
}
