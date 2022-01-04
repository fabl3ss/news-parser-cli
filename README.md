# news-parser-cli
Simple CLI which allows you to receive news depending on the parameters passed to it

![alt text](https://github.com/fabl3ss/news-parser-cli/blob/main/src/cli/img/news_parser_green.png?raw=true)

# HOW TO USE?
**Query structure:**\
./exe_name get -lang <*string*> -country <*string*> -category <*string*> -kwords <*string*>

**Parameters:**\
  -lang (*necessarily*)\
  -country (*optional*)\
  -category (*necessarily if no -kwords arguments*)\
  -kwords (*necessarily if no -category arguments*)

**Examples:**\
./web-aggregator get -lang "en" -kwords "usa"\
./web-aggregator get -lang "ru" -kwords "биткоин майнинг"\
./web-aggregator get -lang "ru" -country "ua" -category "entertainment"
                                                    
**Response example:**\
{\
    "status": "success",\
    "totalResults": 1,\
    "results": [\
        {\
            "title": "Холли Берри вышла замуж за музыканта на тропическом острове",\
            "link": "https://www.mk.ru/culture/2022/01/03/kholli-berri-vyshla-zamuzh-za-muzykanta-na-tropicheskom-ostrove.html",\
            "keywords": [\
                "Культура"\
            ],\
            "creator": null,\
            "video_url": null,\
            "description": "55-летняя американская актриса, режиссер, лауреат премий \"Оскар\", \"Золотой глобус\" и \"Эмми! Холли Берри разместила на своей ст
ранице в Instagram фото со свадебной церемонии",\
            "content": null,\
            "pubDate": "2022-01-02 21:43:37",\
            "full_description": "55-летняя американская актриса, режиссер, лауреат премий \"Оскар\", \"Золотой глобус\" и \"Эмми! Холли Берри разместила на сво
ей\ странице в Instagram фото со свадебной церемонии. Женихом актрисы стал ее бойфренд, 52-летний певец Ван Хант, с которым она состоит в отношениях с осени 202
0 года.\ \"Что ж, Теперь все официально\", - подписала романтический снимок Берри. Впрочем, подписчики не спешат поздравлять актрису с четвертым бракосочетанием
,\ так как не исключают, что пост был просто розыгрышем.",\
            "image_url": "https://static.mk.ru/upload/entities/2022/01/03/00/articles/detailPicture/fb/72/ec/7c/3d0ffb695a4672325ff8fa6798e02209.jpg",\
            "source_id": "mk"\
        }\
    ],\
    "nextPage": null\
}

\
**All possible countries:**
* Argentina - ar
* Australia - au
* Austria - at
* Belgium - be
* Brazil - br
* Bulgaria - bg
* Canada - ca
* China - cn
* Colombia - co
* Cuba - cu
* Czech republic - cz
* Egypt - eg
* France - fr
* Germany - de
* Greece - gr
* Hong kong - hk
* Hungary - hu
* India - in
* Indonesia - id
* Ireland - ie
* Israel - il
* Italy - it
* Japan - jp
* Latvia - lv
* Lebanon - lb
* Lithuania - lt
* Malaysia - my
* Mexico - mx
* Morocco - ma
* Netherland - nl
* New zealand - nz
* Nigeria - ng
* North korea - kp
* Norway - no
* Pakistan - pk
* Philippines - ph
* Poland - pl
* Portugal - pt
* Romania - ro
* Russia - ru
* Saudi arabia - sa
* Serbia - rs
* Singapore - sg
* Slovakia - sk
* Slovenia - si
* South africa - za
* South korea - kr
* Spain - es
* Sweden - se
* Switzerland - ch
* Taiwan - tw
* Thailand - th
* Turkey - tr
* Ukraine - ua
* United arab emirates - ae
* United kingdom - gb
* United states of america - us
* Venezuela - ve
  
  
**All possible categories:**
* business
* entertainment
* environment
* food
* health
* politics
* science
* sports
* technology
* top
* world

**All possible languages:**
* Arabic - ar
* Bosnian - bs
* Bulgarian - bg
* Chinese - zh
* Croatian - hr
* Czech - cs
* Dutch - nl
* English - en
* French - fr
* German - de
* Greek - el
* Hebrew - he
* Hindi - hi
* Hungarian - hu
* Indonesian - in
* Italian - it
* Japanese - jp
* Korean - ko
* Latvian - lv
* Lithuanian - lt
* Malay - ms
* Norwegian - no
* Polish - pl
* Portuguese - pt
* Romanian - ro
* Russian - ru
* Serbian - sr
* Slovak - sk
* Slovenian - sl
* Spanish - es
* Swedish - sv
* Thai - th
* Turkish - tr
* Ukrainian - uk
