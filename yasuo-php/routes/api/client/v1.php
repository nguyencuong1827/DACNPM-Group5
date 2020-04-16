<?php

Route::group(['prefix' => 'me'], function () {
    Route::post('/', 'MeController@logout');
});
