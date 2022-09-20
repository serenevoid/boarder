package models

import (
	"testing"
)

func Test_Get_threads_from_json(t *testing.T) {
    data := "[{\"page\":1,\"threads\":[{\"no\":6872254,\"last_modified\":1489266886,\"replies\":4}]}]"
    thread := Get_threads_from_json([]byte(data))
    if thread[0] != 6872254 {
        t.Error()
    }
}

func Test_Get_posts_from_json(t *testing.T) {
    data := "{\"posts\":[{\"no\":6872254,\"sticky\":1,\"closed\":1,\"now\":\"03\\/11\\/17(Sat)16:09:30\",\"name\":\"The \\/wg\\/ Bros\",\"sub\":\"Welcome to \\/wg\\/ - Wallpapers\\/General\",\"com\":\"New to \\/wg\\/? Lets get you started.<br>~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~<br>1) Look before you post<br>2) Post more than one, sharing is caring<br>3) We already have an Image Modification thread<br>~ ~ (ALL colorsplash, watermarks, photoshop requests)<br>4) We already have a Desktop thread<br>~ ~ (ALL desktops, rating, and theme\\/hax questions)<br>5) Share anything WP related!<br>~ ~ (NO low res\\/quality, illegal content;<br>~ ~ ~ anime goes in \\/w\\/, this is not \\/r\\/)<br><br>We on \\/wg\\/ love WPs and we love sharing them.<br>~ ~ ~ That&#039;s why we&#039;re here.<br>~ ~ ~ ~ ~ Now lets do it!\",\"filename\":\"stickyop\",\"ext\":\".jpg\",\"w\":840,\"h\":672,\"tn_w\":250,\"tn_h\":200,\"tim\":1489266570954,\"time\":1489266570,\"md5\":\"n2gHJOSc1QFSRNUjlcybBw==\",\"fsize\":224079,\"resto\":0,\"semantic_url\":\"welcome-to-wg-wallpapersgeneral\",\"replies\":4,\"images\":3,\"unique_ips\":1}]}"
    post := Get_posts_from_json([]byte(data))
    if post[0].No != 6872254 {
        t.Error()
    }
}
