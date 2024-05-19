package api

var (
	Props = `
    {
        "props": {
          "pageProps": {
            "user": {
              "id": "user-ieBaO6Zyq6qmnqo6k4qEyj9Q",
              "name": "singmertikahi@mail.com",
              "email": "singmertikahi@mail.com",
              "image": "https://s.gravatar.com/avatar/4e8ecb3eecfe9e87005c66f1af7d18df?s=480\u0026r=pg\u0026d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsi.png",
              "picture": "https://s.gravatar.com/avatar/4e8ecb3eecfe9e87005c66f1af7d18df?s=480\u0026r=pg\u0026d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsi.png",
              "idp": "auth0",
              "iat": 1716055529,
              "mfa": false,
              "groups": [],
              "intercom_hash": "39c96562fa85b492d623bd9caf4826f5fd365f7792c591dac0fc6164f80f6196"
            },
            "serviceStatus": {},
            "serviceAnnouncement": { "public": {}, "paid": {} },
            "userCountry": "US",
            "userRegion": "New York",
            "userRegionCode": "NY",
            "cfConnectingIp": "23.95.156.48",
            "cfIpCity": "Buffalo",
            "cfTimezone": "America/New_York",
            "userLocale": "en",
            "statsig": {
              "payload": {
                "feature_gates": {
                  "29898433": {
                    "name": "29898433",
                    "value": true,
                    "rule_id": "5hlR2AjP7h2nJXOxZHKzKH:100.00:2",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "48895201": {
                    "name": "48895201",
                    "value": false,
                    "rule_id": "3i3FRD5a2bDACgRX2SCksw",
                    "secondary_exposures": []
                  },
                  "73891701": {
                    "name": "73891701",
                    "value": true,
                    "rule_id": "4M51I9cMzYqCiiDU2N0eXN:100.00:1",
                    "secondary_exposures": [
                      {
                        "gate": "segment:chatgpt_unpaid_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:chatgpt_paid_users",
                        "gateValue": "true",
                        "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                      }
                    ]
                  },
                  "81143276": {
                    "name": "81143276",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "87687712": {
                    "name": "87687712",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "174366048": {
                    "name": "174366048",
                    "value": true,
                    "rule_id": "bhPM7FsN2H1vnBUrxrg6v:100.00:3",
                    "secondary_exposures": [
                      {
                        "gate": "segment:chatgpt_voice_team",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:friends_and_family_nda",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:chatgpt_paid_users",
                        "gateValue": "true",
                        "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                      },
                      {
                        "gate": "chatgpt-web-voice-message-playback",
                        "gateValue": "true",
                        "ruleID": "6VUF6Z1JaUKZF7RS6uSjUu:100.00:6"
                      }
                    ]
                  },
                  "211801712": {
                    "name": "211801712",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "226799327": {
                    "name": "226799327",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "437245079": {
                    "name": "437245079",
                    "value": true,
                    "rule_id": "1KKPtXmOjxZamlgM2DzXlg",
                    "secondary_exposures": []
                  },
                  "458009956": {
                    "name": "458009956",
                    "value": true,
                    "rule_id": "6LgEwykI8hHnFAQ5EkKTX2:100.00:1",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_free_preview_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_paid_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "544659310": {
                    "name": "544659310",
                    "value": true,
                    "rule_id": "1Opz1br1XsEM6Uajp6F1XU:100.00:3",
                    "secondary_exposures": []
                  },
                  "607723475": {
                    "name": "607723475",
                    "value": false,
                    "rule_id": "2e0cON2LQDfCjkZQ6yzHg5",
                    "secondary_exposures": []
                  },
                  "716475478": {
                    "name": "716475478",
                    "value": true,
                    "rule_id": "48mToGJmOtQ5aRx9MCjOor",
                    "secondary_exposures": []
                  },
                  "716722001": {
                    "name": "716722001",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "730579107": {
                    "name": "730579107",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "740954505": {
                    "name": "740954505",
                    "value": true,
                    "rule_id": "2Ey97LJY7QymPfJfxULJwr:85.00:5",
                    "secondary_exposures": [
                      {
                        "gate": "segment:anonymous_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "internal-employee-only-chatgpt-team",
                        "gateValue": "false",
                        "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                      },
                      {
                        "gate": "segment:enterprise",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_free_preview_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_paid_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "771993664": {
                    "name": "771993664",
                    "value": false,
                    "rule_id": "1dDdxrXy1KadHtYusuGRgR:0.00:5",
                    "secondary_exposures": []
                  },
                  "857491970": {
                    "name": "857491970",
                    "value": true,
                    "rule_id": "2LpyNTIfbYC4R6XS6dfdIL",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "859327441": {
                    "name": "859327441",
                    "value": false,
                    "rule_id": "2kLMbBkcSZ8TkzHgzgbyKV",
                    "secondary_exposures": []
                  },
                  "874794369": {
                    "name": "874794369",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "1026575464": {
                    "name": "1026575464",
                    "value": true,
                    "rule_id": "1f4cSeenpcmWQ9eMKKQhF5",
                    "secondary_exposures": []
                  },
                  "1056591003": {
                    "name": "1056591003",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:internal_oai_gptv_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:chatgpt_unpaid_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-mainline-model-transition-unpaid-sahara-v",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "1070272635": {
                    "name": "1070272635",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:friends_and_family_nda",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:feather_memory_trainers",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "1123641887": {
                    "name": "1123641887",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt_anon_chat_enabled",
                        "gateValue": "true",
                        "ruleID": "22nVhoL17eyMvGWgFrDfZe"
                      }
                    ]
                  },
                  "1123642877": {
                    "name": "1123642877",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt_anon_chat_enabled",
                        "gateValue": "true",
                        "ruleID": "22nVhoL17eyMvGWgFrDfZe"
                      }
                    ]
                  },
                  "1162285750": {
                    "name": "1162285750",
                    "value": false,
                    "rule_id": "5SD2EhDYytn75oTAjOCSiL",
                    "secondary_exposures": []
                  },
                  "1198030896": {
                    "name": "1198030896",
                    "value": true,
                    "rule_id": "2aaGr86KHS6GUSs6I9F2cV",
                    "secondary_exposures": []
                  },
                  "1200224890": {
                    "name": "1200224890",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:anonymous_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:enterprise",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "internal-employee-only-chatgpt-team",
                        "gateValue": "false",
                        "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                      },
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_free_preview_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_paid_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt_fruit_juice",
                        "gateValue": "true",
                        "ruleID": "2Ey97LJY7QymPfJfxULJwr:85.00:5"
                      }
                    ]
                  },
                  "1306782371": {
                    "name": "1306782371",
                    "value": true,
                    "rule_id": "1zW3Cog0vZrd32oA1luIPz",
                    "secondary_exposures": []
                  },
                  "1426009137": {
                    "name": "1426009137",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:chatgpt_voice_team",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "1428536080": {
                    "name": "1428536080",
                    "value": false,
                    "rule_id": "2lFzVsatjvpp5kR8PhJaEl",
                    "secondary_exposures": []
                  },
                  "1568195116": {
                    "name": "1568195116",
                    "value": false,
                    "rule_id": "KKhehNjRxNHoIWKNkqpU2",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "1654633603": {
                    "name": "1654633603",
                    "value": true,
                    "rule_id": "2WWYf2mhNtJ9juDjcM3gmE",
                    "secondary_exposures": [
                      {
                        "gate": "segment:anonymous_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "internal-employee-only-chatgpt-team",
                        "gateValue": "false",
                        "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                      },
                      {
                        "gate": "segment:enterprise",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_free_preview_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_paid_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt_fruit_juice",
                        "gateValue": "true",
                        "ruleID": "2Ey97LJY7QymPfJfxULJwr:85.00:5"
                      }
                    ]
                  },
                  "1703773460": {
                    "name": "1703773460",
                    "value": false,
                    "rule_id": "63qNN5LSdeNXfSPo2Mwuv2:0.00:2",
                    "secondary_exposures": []
                  },
                  "1820151945": {
                    "name": "1820151945",
                    "value": true,
                    "rule_id": "3dG8z8CNFtgpXJNyOi8ijR",
                    "secondary_exposures": [
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:model_switcher_team",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:chatgpt_paid_users",
                        "gateValue": "true",
                        "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                      }
                    ]
                  },
                  "1825130190": {
                    "name": "1825130190",
                    "value": true,
                    "rule_id": "Nef2uMceNUF9U3ZYwSbpD",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "1828361965": {
                    "name": "1828361965",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "1852364556": {
                    "name": "1852364556",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "1880477506": {
                    "name": "1880477506",
                    "value": true,
                    "rule_id": "7fFsLeJu6pyCq9tRWmX9em:100.00:6",
                    "secondary_exposures": []
                  },
                  "1883608273": {
                    "name": "1883608273",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "1923022511": {
                    "name": "1923022511",
                    "value": true,
                    "rule_id": "6VUF6Z1JaUKZF7RS6uSjUu:100.00:6",
                    "secondary_exposures": [
                      {
                        "gate": "segment:chatgpt_voice_team",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:friends_and_family_nda",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:chatgpt_paid_users",
                        "gateValue": "true",
                        "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                      }
                    ]
                  },
                  "1938289220": {
                    "name": "1938289220",
                    "value": false,
                    "rule_id": "1vWEhLfJ39edDu8ZjE19dl:50.00:4",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:contractors",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:sidekick_windows_contractors",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:friends_and_family_nda",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:enterprise",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:chatgpt_paid_users",
                        "gateValue": "true",
                        "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                      }
                    ]
                  },
                  "2000076788": {
                    "name": "2000076788",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:contractors",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2004295676": {
                    "name": "2004295676",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:sonic-testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-sonic-internal",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-search-demo-ui",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2046333385": {
                    "name": "2046333385",
                    "value": false,
                    "rule_id": "4nM2ehmgoDQIv69B0zohb6",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2166885617": {
                    "name": "2166885617",
                    "value": false,
                    "rule_id": "1ORnWBMVyZ42eAp03fe5Nr",
                    "secondary_exposures": []
                  },
                  "2173548801": {
                    "name": "2173548801",
                    "value": true,
                    "rule_id": "22nVhoL17eyMvGWgFrDfZe",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2438078654": {
                    "name": "2438078654",
                    "value": true,
                    "rule_id": "VNBdYfNSNNUyDdMpaiaZ9:100.00:3",
                    "secondary_exposures": []
                  },
                  "2449630478": {
                    "name": "2449630478",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "2542814321": {
                    "name": "2542814321",
                    "value": false,
                    "rule_id": "6Usmt3Th0mkUkVQJ0oyeST:0.00:1",
                    "secondary_exposures": []
                  },
                  "2546838099": {
                    "name": "2546838099",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:sonic-testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-sonic-internal",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2562876640": {
                    "name": "2562876640",
                    "value": true,
                    "rule_id": "326czTZeZ0RX0ypR0c5Bb6:100.00:15",
                    "secondary_exposures": []
                  },
                  "2689530298": {
                    "name": "2689530298",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2865702999": {
                    "name": "2865702999",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2865703032": {
                    "name": "2865703032",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2893926039": {
                    "name": "2893926039",
                    "value": false,
                    "rule_id": "1XOa3bw1GRog1lbmWKnaF9",
                    "secondary_exposures": []
                  },
                  "2947776398": {
                    "name": "2947776398",
                    "value": false,
                    "rule_id": "2YzNwxxdROEWn1lPwJbYtr",
                    "secondary_exposures": []
                  },
                  "3011278998": {
                    "name": "3011278998",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "3060586957": {
                    "name": "3060586957",
                    "value": false,
                    "rule_id": "1agh2VKHmwqbsORiGxd3v5",
                    "secondary_exposures": []
                  },
                  "3134329717": {
                    "name": "3134329717",
                    "value": false,
                    "rule_id": "2UWRV6bObClvvj8jLkP7kJ",
                    "secondary_exposures": []
                  },
                  "3254335143": {
                    "name": "3254335143",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": []
                  },
                  "3320547351": {
                    "name": "3320547351",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "3376455464": {
                    "name": "3376455464",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "3455796462": {
                    "name": "3455796462",
                    "value": true,
                    "rule_id": "YbNn9kuhmMn3DdnzEDc6G",
                    "secondary_exposures": []
                  },
                  "3483957186": {
                    "name": "3483957186",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:anonymous_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "3501550327": {
                    "name": "3501550327",
                    "value": false,
                    "rule_id": "HEU9zVkuwUfmXHhGcckPQ",
                    "secondary_exposures": []
                  },
                  "3511639124": {
                    "name": "3511639124",
                    "value": true,
                    "rule_id": "1y9t5DBS0nz9mqbQRCrmCP:100.00:5",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:plus_qa_-_gizmos",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:store_dogfooding",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:anonymous_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:gizmo_store_holdout_exceptions",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "3643162575": {
                    "name": "3643162575",
                    "value": true,
                    "rule_id": "56j9Tu0CGzcoiVbnTxuR7O",
                    "secondary_exposures": []
                  },
                  "3700615661": {
                    "name": "3700615661",
                    "value": true,
                    "rule_id": "66covjaoZoe9pQR4I68jOB",
                    "secondary_exposures": []
                  },
                  "3912844760": {
                    "name": "3912844760",
                    "value": false,
                    "rule_id": "default",
                    "secondary_exposures": [
                      {
                        "gate": "segment:european_legal_restricted",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "3922145230": {
                    "name": "3922145230",
                    "value": false,
                    "rule_id": "14DZA2LumaPqAdCo52CrUB",
                    "secondary_exposures": [
                      {
                        "gate": "segment:internal_ips",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "3941367332": {
                    "name": "3941367332",
                    "value": true,
                    "rule_id": "3tesvDaP0kRy0JAXNdSJJn:100.00:3",
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                        "gateValue": "true",
                        "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                      }
                    ]
                  },
                  "4226692983": {
                    "name": "4226692983",
                    "value": true,
                    "rule_id": "6sEu91zwlBGSKOqFiNpGlA:100.00:2",
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:enterprise",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:teams_plan",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  }
                },
                "dynamic_configs": {
                  "75342665": {
                    "name": "75342665",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": ["is_anon_chat_enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "75342698": {
                    "name": "75342698",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": ["is_anon_chat_enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "101132775": {
                    "name": "101132775",
                    "value": { "show_new_ui": true },
                    "group": "launchedGroup",
                    "rule_id": "launchedGroup",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false
                  },
                  "171317775": {
                    "name": "171317775",
                    "value": {},
                    "group": "prestart",
                    "rule_id": "prestart",
                    "is_device_based": true,
                    "secondary_exposures": [],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false
                  },
                  "340466392": {
                    "name": "340466392",
                    "value": {
                      "google.com": [
                        "google.com",
                        "googleapis.com",
                        "www.googleapis.com",
                        "oauth2.googleapis.com",
                        "cloudfunctions.net"
                      ],
                      "microsoftonline.com": [
                        "microsoftonline.com",
                        "azurewebsites.net",
                        "graph.microsoft.com",
                        "microsoft.com",
                        "azure.com",
                        "sinequa.com"
                      ],
                      "adobelogin.com": ["adobe.com", "adobelogin.com"],
                      "useflorian.com": ["useflorian.com", "convex.site"],
                      "yahoo.com": ["yahooapis.com", "yahoo.com"],
                      "visualstudio.com": ["visualstudio.com", "azure.com"],
                      "atlassian.com": ["atlassian.net", "atlassian.com"]
                    },
                    "group": "default",
                    "rule_id": "default",
                    "is_device_based": true,
                    "secondary_exposures": []
                  },
                  "445876862": {
                    "name": "445876862",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": [
                      "is_anon_chat_enabled",
                      "is_no_auth_welcome_modal_enabled",
                      "no_auth_soft_rate_limit",
                      "no_auth_hard_rate_limit",
                      "should_show_no_auth_signup_banner"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "855925714": {
                    "name": "855925714",
                    "value": {
                      "should_simplify_modal": false,
                      "is_simplified_sharing_modal_enabled": false,
                      "is_social_share_options_enabled": false,
                      "is_update_shared_links_enabled": false,
                      "is_discoverability_toggle_enabled": false,
                      "show_copylink_state_if_no_updates": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_simplified_sharing_modal_enabled",
                      "is_social_share_options_enabled",
                      "is_update_shared_links_enabled",
                      "is_discoverability_toggle_enabled"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "889327355": {
                    "name": "889327355",
                    "value": { "is-enabled": true },
                    "group": "layerAssignment",
                    "rule_id": "layerAssignment",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": ["is-enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": true,
                    "is_in_layer": true
                  },
                  "1193714591": {
                    "name": "1193714591",
                    "value": {
                      "is_dynamic_model_enabled": false,
                      "show_message_model_info": false,
                      "show_message_regenerate_model_selector": true,
                      "is_conversation_model_switching_allowed": true,
                      "show_rate_limit_downgrade_banner": true,
                      "config": {
                        "impl": "rm_score",
                        "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                        "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                      },
                      "show_message_regenerate_model_selector_on_every_message": true,
                      "is_AG8PqS2q_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_dynamic_model_enabled",
                      "show_message_regenerate_model_selector",
                      "is_conversation_model_switching_allowed",
                      "show_rate_limit_downgrade_banner"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "1193714592": {
                    "name": "1193714592",
                    "value": {
                      "is_dynamic_model_enabled": false,
                      "show_message_model_info": false,
                      "show_message_regenerate_model_selector": true,
                      "is_conversation_model_switching_allowed": true,
                      "show_rate_limit_downgrade_banner": true,
                      "config": {
                        "impl": "rm_score",
                        "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                        "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                      },
                      "show_message_regenerate_model_selector_on_every_message": true,
                      "is_AG8PqS2q_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_dynamic_model_enabled",
                      "show_message_regenerate_model_selector",
                      "is_conversation_model_switching_allowed",
                      "show_rate_limit_downgrade_banner",
                      "config"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "1193714593": {
                    "name": "1193714593",
                    "value": {
                      "is_dynamic_model_enabled": false,
                      "show_message_model_info": false,
                      "show_message_regenerate_model_selector": true,
                      "is_conversation_model_switching_allowed": true,
                      "show_rate_limit_downgrade_banner": true,
                      "config": {
                        "impl": "rm_score",
                        "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                        "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                      },
                      "show_message_regenerate_model_selector_on_every_message": true,
                      "is_AG8PqS2q_enabled": false
                    },
                    "group": "prestart",
                    "rule_id": "prestart",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_dynamic_model_enabled",
                      "show_message_regenerate_model_selector",
                      "is_conversation_model_switching_allowed",
                      "show_rate_limit_downgrade_banner",
                      "config",
                      "show_message_regenerate_model_selector_on_every_message"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "1237904433": {
                    "name": "1237904433",
                    "value": { "is-enabled": true },
                    "group": "launchedGroup",
                    "rule_id": "launchedGroup",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": ["is-enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "1277879515": {
                    "name": "1277879515",
                    "value": { "show_ui": true },
                    "group": "launchedGroup",
                    "rule_id": "launchedGroup",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "segment:friends_and_family_nda",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:feather_memory_trainers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-allow-memory-use-admin",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false
                  },
                  "1312738797": {
                    "name": "1312738797",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "launchedGroup",
                    "rule_id": "launchedGroup",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": ["anon-is-spanish-translation-enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "1476759596": {
                    "name": "1476759596",
                    "value": {},
                    "group": "layerAssignment",
                    "rule_id": "layerAssignment",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "is_user_in_experiment": false,
                    "is_experiment_active": true
                  },
                  "1547924522": {
                    "name": "1547924522",
                    "value": { "show_ui": true, "enabled": true },
                    "group": "launchedGroup",
                    "rule_id": "launchedGroup",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt-paid-users",
                        "gateValue": "true",
                        "ruleID": "7stvavArTuHK8XvexxjCiu"
                      }
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false
                  },
                  "1887692490": {
                    "name": "1887692490",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": [
                      "is_anon_chat_enabled",
                      "is_try_it_first_on_login_page_enabled"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "1908072088": {
                    "name": "1908072088",
                    "value": {
                      "blocking_modal": true,
                      "single_button_variant": true,
                      "dropdown_tooltip": false,
                      "conversation_model_switcher_badge": false
                    },
                    "group": "default",
                    "rule_id": "default",
                    "is_device_based": false,
                    "secondary_exposures": []
                  },
                  "2112835930": {
                    "name": "2112835930",
                    "value": {},
                    "group": "targetingGate",
                    "rule_id": "targetingGate",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "segment:anonymous_users",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:enterprise",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "internal-employee-only-chatgpt-team",
                        "gateValue": "false",
                        "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                      },
                      {
                        "gate": "segment:plusqa_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_free_preview_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:scallion_paid_testers",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt_fruit_juice",
                        "gateValue": "true",
                        "ruleID": "2Ey97LJY7QymPfJfxULJwr:85.00:5"
                      },
                      {
                        "gate": "chatgpt_fruit_juice_inverse",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": true
                  },
                  "2144828235": {
                    "name": "2144828235",
                    "value": {},
                    "group": "layerAssignment",
                    "rule_id": "layerAssignment",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "is_user_in_experiment": false,
                    "is_experiment_active": true
                  },
                  "2176179500": {
                    "name": "2176179500",
                    "value": { "should_enable_zh_tw": false },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": ["should_enable_zh_tw"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2331692225": {
                    "name": "2331692225",
                    "value": {},
                    "group": "default",
                    "rule_id": "default",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-browse-sahara",
                        "gateValue": "false",
                        "ruleID": "default"
                      }
                    ]
                  },
                  "2370410632": {
                    "name": "2370410632",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": ["is_anon_chat_enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2370410634": {
                    "name": "2370410634",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": ["is_anon_chat_enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2575540972": {
                    "name": "2575540972",
                    "value": {
                      "is_dynamic_model_enabled": false,
                      "show_message_model_info": false,
                      "show_message_regenerate_model_selector": true,
                      "is_conversation_model_switching_allowed": true,
                      "show_rate_limit_downgrade_banner": true,
                      "config": {
                        "impl": "rm_score",
                        "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                        "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                      },
                      "show_message_regenerate_model_selector_on_every_message": true,
                      "is_AG8PqS2q_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_dynamic_model_enabled",
                      "show_message_regenerate_model_selector",
                      "is_conversation_model_switching_allowed",
                      "show_rate_limit_downgrade_banner"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2576123431": {
                    "name": "2576123431",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": [
                      "is_anon_chat_enabled",
                      "is_anon_chat_enabled_for_new_users_only"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2635734272": {
                    "name": "2635734272",
                    "value": {
                      "should_simplify_modal": false,
                      "is_simplified_sharing_modal_enabled": false,
                      "is_social_share_options_enabled": false,
                      "is_update_shared_links_enabled": false,
                      "is_discoverability_toggle_enabled": false,
                      "show_copylink_state_if_no_updates": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_simplified_sharing_modal_enabled",
                      "is_social_share_options_enabled",
                      "is_update_shared_links_enabled",
                      "is_discoverability_toggle_enabled"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2745767556": {
                    "name": "2745767556",
                    "value": {
                      "is_team_enabled": true,
                      "is_yearly_plus_subscription_enabled": false,
                      "is_split_between_personal_and_business_enabled": false
                    },
                    "group": "prestart",
                    "rule_id": "prestart",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      }
                    ],
                    "explicit_parameters": [
                      "is_split_between_personal_and_business_enabled"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2811617088": {
                    "name": "2811617088",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": [
                      "is_anon_chat_enabled",
                      "is_no_auth_welcome_modal_enabled",
                      "no_auth_soft_rate_limit",
                      "no_auth_hard_rate_limit",
                      "should_show_no_auth_signup_banner"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "2891383514": {
                    "name": "2891383514",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "targetingGate",
                    "rule_id": "targetingGate",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      },
                      {
                        "gate": "chatgpt_no_auth_thanks_for_trying_20240515_gate",
                        "gateValue": "false",
                        "ruleID": "2lFzVsatjvpp5kR8PhJaEl"
                      }
                    ],
                    "explicit_parameters": [
                      "no_auth_soft_rate_limit",
                      "is_no_auth_soft_rate_limit_modal_enabled"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": true,
                    "is_in_layer": true
                  },
                  "2919366999": {
                    "name": "2919366999",
                    "value": {
                      "should_simplify_modal": false,
                      "is_simplified_sharing_modal_enabled": false,
                      "is_social_share_options_enabled": false,
                      "is_update_shared_links_enabled": false,
                      "is_discoverability_toggle_enabled": false,
                      "show_copylink_state_if_no_updates": false
                    },
                    "group": "3fdACLHPKSOwzpRi35OaGG",
                    "rule_id": "3fdACLHPKSOwzpRi35OaGG",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                        "gateValue": "true",
                        "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                      }
                    ],
                    "explicit_parameters": ["show_copylink_state_if_no_updates"],
                    "is_user_in_experiment": true,
                    "is_experiment_active": true,
                    "is_in_layer": true
                  },
                  "2980740905": {
                    "name": "2980740905",
                    "value": {
                      "is_team_enabled": true,
                      "is_yearly_plus_subscription_enabled": false,
                      "is_split_between_personal_and_business_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      }
                    ],
                    "explicit_parameters": ["is_yearly_plus_subscription_enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "3168246095": {
                    "name": "3168246095",
                    "value": { "gizmo_ids": [] },
                    "group": "default",
                    "rule_id": "default",
                    "is_device_based": false,
                    "secondary_exposures": []
                  },
                  "3299284219": {
                    "name": "3299284219",
                    "value": {
                      "is_dynamic_model_enabled": false,
                      "show_message_model_info": false,
                      "show_message_regenerate_model_selector": true,
                      "is_conversation_model_switching_allowed": true,
                      "show_rate_limit_downgrade_banner": true,
                      "config": {
                        "impl": "rm_score",
                        "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                        "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                      },
                      "show_message_regenerate_model_selector_on_every_message": true,
                      "is_AG8PqS2q_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_dynamic_model_enabled",
                      "show_message_regenerate_model_selector",
                      "is_conversation_model_switching_allowed",
                      "show_rate_limit_downgrade_banner",
                      "show_message_regenerate_model_selector_on_every_message"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "3340924940": {
                    "name": "3340924940",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": [
                      "is_anon_chat_enabled",
                      "should_show_anon_login_header_on_desktop"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "3414158273": {
                    "name": "3414158273",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "launchedGroup",
                    "rule_id": "launchedGroup",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": [
                      "is_try_it_first_on_login_page_enabled",
                      "is_no_auth_welcome_back_modal_enabled"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "3414159263": {
                    "name": "3414159263",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "layerAssignment",
                    "rule_id": "layerAssignment",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ],
                    "explicit_parameters": ["is_no_auth_welcome_back_modal_enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": true,
                    "is_in_layer": true
                  },
                  "3758098899": {
                    "name": "3758098899",
                    "value": {
                      "is_team_enabled": true,
                      "is_yearly_plus_subscription_enabled": false,
                      "is_split_between_personal_and_business_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      }
                    ],
                    "explicit_parameters": ["is_team_enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "3758099832": {
                    "name": "3758099832",
                    "value": {},
                    "group": "prestart",
                    "rule_id": "prestart",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false
                  },
                  "3917128856": {
                    "name": "3917128856",
                    "value": {
                      "is_dynamic_model_enabled": false,
                      "show_message_model_info": false,
                      "show_message_regenerate_model_selector": true,
                      "is_conversation_model_switching_allowed": true,
                      "show_rate_limit_downgrade_banner": true,
                      "config": {
                        "impl": "rm_score",
                        "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                        "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                      },
                      "show_message_regenerate_model_selector_on_every_message": true,
                      "is_AG8PqS2q_enabled": false
                    },
                    "group": "abandoned",
                    "rule_id": "abandoned",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [
                      "is_dynamic_model_enabled",
                      "show_message_regenerate_model_selector",
                      "is_conversation_model_switching_allowed",
                      "show_rate_limit_downgrade_banner",
                      "config"
                    ],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  },
                  "4204547666": {
                    "name": "4204547666",
                    "value": { "is-enabled": true },
                    "group": "Nzc6Xnht6tIVmb48Ejg1T:override",
                    "rule_id": "Nzc6Xnht6tIVmb48Ejg1T:override",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "segment:internal_ips",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-localization-preview",
                        "gateValue": "false",
                        "ruleID": "14DZA2LumaPqAdCo52CrUB"
                      },
                      {
                        "gate": "chatgpt-localization-allowlist",
                        "gateValue": "true",
                        "ruleID": "66covjaoZoe9pQR4I68jOB"
                      }
                    ],
                    "explicit_parameters": ["is-enabled"],
                    "is_user_in_experiment": false,
                    "is_experiment_active": false,
                    "is_in_layer": true
                  }
                },
                "layer_configs": {
                  "468168202": {
                    "name": "468168202",
                    "value": {
                      "is_team_enabled": true,
                      "is_yearly_plus_subscription_enabled": false,
                      "is_split_between_personal_and_business_enabled": false
                    },
                    "group": "default",
                    "rule_id": "default",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      }
                    ],
                    "explicit_parameters": [],
                    "undelegated_secondary_exposures": [
                      {
                        "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      }
                    ]
                  },
                  "1238742812": {
                    "name": "1238742812",
                    "value": { "should_enable_zh_tw": false },
                    "group": "default",
                    "rule_id": "default",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [],
                    "undelegated_secondary_exposures": []
                  },
                  "1547743984": {
                    "name": "1547743984",
                    "value": {
                      "should_simplify_modal": false,
                      "is_simplified_sharing_modal_enabled": false,
                      "is_social_share_options_enabled": false,
                      "is_update_shared_links_enabled": false,
                      "is_discoverability_toggle_enabled": false,
                      "show_copylink_state_if_no_updates": false
                    },
                    "group": "3fdACLHPKSOwzpRi35OaGG",
                    "rule_id": "3fdACLHPKSOwzpRi35OaGG",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                        "gateValue": "true",
                        "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                      }
                    ],
                    "explicit_parameters": ["show_copylink_state_if_no_updates"],
                    "allocated_experiment_name": "2919366999",
                    "is_experiment_active": true,
                    "is_user_in_experiment": true,
                    "undelegated_secondary_exposures": []
                  },
                  "2723963139": {
                    "name": "2723963139",
                    "value": {
                      "is_dynamic_model_enabled": false,
                      "show_message_model_info": false,
                      "show_message_regenerate_model_selector": true,
                      "is_conversation_model_switching_allowed": true,
                      "show_rate_limit_downgrade_banner": true,
                      "config": {
                        "impl": "rm_score",
                        "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                        "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                      },
                      "show_message_regenerate_model_selector_on_every_message": true,
                      "is_AG8PqS2q_enabled": false
                    },
                    "group": "default",
                    "rule_id": "default",
                    "is_device_based": false,
                    "secondary_exposures": [],
                    "explicit_parameters": [],
                    "undelegated_secondary_exposures": []
                  },
                  "3048336830": {
                    "name": "3048336830",
                    "value": { "is-enabled": true },
                    "group": "Nzc6Xnht6tIVmb48Ejg1T:override",
                    "rule_id": "Nzc6Xnht6tIVmb48Ejg1T:override",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "segment:internal_ips",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-localization-preview",
                        "gateValue": "false",
                        "ruleID": "14DZA2LumaPqAdCo52CrUB"
                      },
                      {
                        "gate": "chatgpt-localization-allowlist",
                        "gateValue": "true",
                        "ruleID": "66covjaoZoe9pQR4I68jOB"
                      }
                    ],
                    "explicit_parameters": ["is-enabled"],
                    "allocated_experiment_name": "4204547666",
                    "is_experiment_active": false,
                    "is_user_in_experiment": false,
                    "undelegated_secondary_exposures": [
                      {
                        "gate": "segment:internal_ips",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "segment:employees",
                        "gateValue": "false",
                        "ruleID": "default"
                      },
                      {
                        "gate": "chatgpt-localization-preview",
                        "gateValue": "false",
                        "ruleID": "14DZA2LumaPqAdCo52CrUB"
                      },
                      {
                        "gate": "chatgpt-localization-allowlist",
                        "gateValue": "true",
                        "ruleID": "66covjaoZoe9pQR4I68jOB"
                      }
                    ]
                  },
                  "3637408529": {
                    "name": "3637408529",
                    "value": {
                      "is_anon_chat_enabled": true,
                      "anon_composer_display_variant": "default",
                      "anon-is-spanish-translation-enabled": true,
                      "should_show_anon_login_header_on_desktop": false,
                      "is_anon_chat_enabled_for_new_users_only": false,
                      "is_try_it_first_on_login_page_enabled": false,
                      "is_no_auth_welcome_modal_enabled": false,
                      "no_auth_soft_rate_limit": 1200,
                      "no_auth_hard_rate_limit": 1200,
                      "should_show_no_auth_signup_banner": true,
                      "is_no_auth_welcome_back_modal_enabled": true,
                      "is_no_auth_soft_rate_limit_modal_enabled": false
                    },
                    "group": "targetingGate",
                    "rule_id": "targetingGate",
                    "is_device_based": false,
                    "secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      },
                      {
                        "gate": "chatgpt_no_auth_thanks_for_trying_20240515_gate",
                        "gateValue": "false",
                        "ruleID": "2lFzVsatjvpp5kR8PhJaEl"
                      }
                    ],
                    "explicit_parameters": [
                      "no_auth_soft_rate_limit",
                      "is_no_auth_soft_rate_limit_modal_enabled"
                    ],
                    "allocated_experiment_name": "2891383514",
                    "is_experiment_active": true,
                    "is_user_in_experiment": false,
                    "undelegated_secondary_exposures": [
                      {
                        "gate": "chatgpt_anon_chat_holdout_20240227",
                        "gateValue": "false",
                        "ruleID": "disabled"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_gate_20240328",
                        "gateValue": "true",
                        "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                      },
                      {
                        "gate": "chatgpt_no_auth_holdout_20240328",
                        "gateValue": "false",
                        "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                      }
                    ]
                  }
                },
                "sdkParams": {},
                "has_updates": true,
                "generator": "statsig-node-sdk",
                "time": 0,
                "evaluated_keys": {
                  "userID": "user-ieBaO6Zyq6qmnqo6k4qEyj9Q",
                  "customIDs": {
                    "WebAnonymousCookieID": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                    "DeviceId": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                    "stableID": "9d8b6e32-4f61-4371-8b15-91823f3016cd"
                  }
                },
                "hash_used": "djb2",
                "user_hash": "2262425412"
              },
              "user": {
                "country": "US",
                "custom": {
                  "has_logged_in_before": true,
                  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:126.0) Gecko/20100101 Firefox/126.0",
                  "is_paid": true,
                  "auth_status": "logged_in",
                  "workspace_id": "53ebfd0f-3d00-4441-ab26-f45a6fd8f2d9"
                },
                "customIDs": {
                  "WebAnonymousCookieID": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                  "DeviceId": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                  "stableID": "9d8b6e32-4f61-4371-8b15-91823f3016cd"
                },
                "ip": "23.95.156.48",
                "locale": "en",
                "privateAttributes": { "email": "singmertikahi@mail.com" },
                "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:126.0) Gecko/20100101 Firefox/126.0",
                "userID": "user-ieBaO6Zyq6qmnqo6k4qEyj9Q"
              }
            },
            "authStatus": "logged_in",
            "serverPrimedAllowBrowserStorageValue": true,
            "showCookieConsentBanner": false,
            "isStorageComplianceEnabled": false,
            "ageVerificationDeadline": null,
            "ageVerificationType": "none",
            "logOutUser": false
          },
          "__N_SSP": true
        },
        "page": "/[[...default]]",
        "query": {},
        "buildId": "-wRE4Obkm_QOW7PLn1x21",
        "assetPrefix": "https://cdn.oaistatic.com",
        "isFallback": false,
        "isExperimentalCompile": false,
        "gssp": true,
        "scriptLoader": []
      }
      
      `
	PropsG = `
      {
          "props": {
            "pageProps": {
              "kind": "chat_page",
              "gizmo": {
                "gizmo": {
                  "id": "g-pmuQfob8d",
                  "organization_id": "org-AWLMos7qkuCxcVO5OhJqPO9X",
                  "short_url": "g-pmuQfob8d-image-generator",
                  "author": {
                    "user_id": "user-xyhelper",
                    "display_name": "NAIF J ALOTAIBI",
                    "link_to": null,
                    "is_verified": true,
                    "selected_display": "name",
                    "will_receive_support_emails": true,
                    "display_socials": []
                  },
                  "voice": { "id": "ember" },
                  "workspace_id": null,
                  "model": null,
                  "instructions": null,
                  "settings": null,
                  "display": {
                    "name": "image generator",
                    "description": "A GPT specialized in generating and refining images with a mix of professional and friendly tone.image generator",
                    "prompt_starters": [
                      "Generate an image of a futuristic city.",
                      "Create a portrait of a fictional character.",
                      "Design a logo for a new tech startup.",
                      "Illustrate a scene from a fantasy novel."
                    ],
                    "profile_pic_id": null,
                    "profile_picture_url": "https://files.oaiusercontent.com/file-M1df4Ab7Ow6QCJ871tBUsi0x?se=2123-11-08T17%3A52%3A06Z\u0026sp=r\u0026sv=2021-08-06\u0026sr=b\u0026rscc=max-age%3D31536000%2C%20immutable\u0026rscd=attachment%3B%20filename%3D40face33-c6ad-4a5d-b402-5f7126e8325f.png\u0026sig=G9qnRNHbnnN1gnz2NzVSyjwWvQ6hrGjr%2Be7hbYgnjRs%3D",
                    "categories": ["dalle"]
                  },
                  "share_recipient": "marketplace",
                  "created_at": "2023-12-02T16:43:09.108969+00:00",
                  "updated_at": "2024-05-18T19:40:27.447894+00:00",
                  "last_interacted_at": null,
                  "num_interactions": null,
                  "tags": [
                    "unreviewable",
                    "public",
                    "reportable",
                    "interactions_disabled"
                  ],
                  "version": null,
                  "live_version": null,
                  "training_disabled": null,
                  "sharing_targets": null,
                  "appeal_info": null,
                  "vanity_metrics": {
                    "num_conversations": null,
                    "num_conversations_str": "6M+",
                    "created_ago_str": "5 months ago",
                    "review_stats": { "total": 638493, "count": 176568 }
                  },
                  "workspace_approval_date": null,
                  "workspace_approved": null,
                  "sharing": null,
                  "current_user_permission": null
                },
                "tools": [
                  {
                    "id": "gzm_cnf_sTkqcaqsAFhjM4b959j5mP5X~gzm_tool_FWZcgGNvH7bOPxkb7MNxvKil",
                    "type": "browser",
                    "settings": null,
                    "metadata": null
                  },
                  {
                    "id": "gzm_cnf_sTkqcaqsAFhjM4b959j5mP5X~gzm_tool_oKrzZ3BvFRBPkJSUiCU3aoEA",
                    "type": "dalle",
                    "settings": null,
                    "metadata": null
                  }
                ],
                "files": [],
                "product_features": {
                  "attachments": {
                    "type": "retrieval",
                    "accepted_mime_types": [
                      "text/x-tex",
                      "text/x-sh",
                      "text/x-typescript",
                      "application/json",
                      "text/x-php",
                      "text/x-csharp",
                      "application/pdf",
                      "text/x-script.python",
                      "text/html",
                      "text/x-c++",
                      "application/vnd.openxmlformats-officedocument.presentationml.presentation",
                      "text/x-ruby",
                      "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
                      "text/markdown",
                      "text/x-c",
                      "application/x-latext",
                      "application/msword",
                      "text/plain",
                      "text/x-java",
                      "text/javascript"
                    ],
                    "image_mime_types": [
                      "image/webp",
                      "image/png",
                      "image/jpeg",
                      "image/gif"
                    ],
                    "can_accept_all_mime_types": true
                  }
                }
              },
              "user": {
                "id": "user-ieBaO6Zyq6qmnqo6k4qEyj9Q",
                "name": "xyhelper@closeai.biz",
                "email": "xyhelper@closeai.biz",
                "image": "https://s.gravatar.com/avatar/4e8ecb3eecfe9e87005c66f1af7d18df?s=480\u0026r=pg\u0026d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsi.png",
                "picture": "https://s.gravatar.com/avatar/4e8ecb3eecfe9e87005c66f1af7d18df?s=480\u0026r=pg\u0026d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fsi.png",
                "idp": "auth0",
                "iat": 1716058929,
                "mfa": false,
                "groups": [],
                "intercom_hash": "39c96562fa85b492d623bd9caf4826f5fd365f7792c591dac0fc6164f80f6196"
              },
              "serviceStatus": {},
              "serviceAnnouncement": { "public": {}, "paid": {} },
              "userCountry": "US",
              "userRegion": "New York",
              "userRegionCode": "NY",
              "cfConnectingIp": "1.1.1.1",
              "cfIpCity": "Buffalo",
              "cfTimezone": "America/New_York",
              "userLocale": "zh-CN",
              "statsig": {
                "payload": {
                  "feature_gates": {
                    "29898433": {
                      "name": "29898433",
                      "value": true,
                      "rule_id": "5hlR2AjP7h2nJXOxZHKzKH:100.00:2",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "48895201": {
                      "name": "48895201",
                      "value": false,
                      "rule_id": "3i3FRD5a2bDACgRX2SCksw",
                      "secondary_exposures": []
                    },
                    "73891701": {
                      "name": "73891701",
                      "value": true,
                      "rule_id": "4M51I9cMzYqCiiDU2N0eXN:100.00:1",
                      "secondary_exposures": [
                        {
                          "gate": "segment:chatgpt_unpaid_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:chatgpt_paid_users",
                          "gateValue": "true",
                          "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                        }
                      ]
                    },
                    "81143276": {
                      "name": "81143276",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "87687712": {
                      "name": "87687712",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "174366048": {
                      "name": "174366048",
                      "value": true,
                      "rule_id": "bhPM7FsN2H1vnBUrxrg6v:100.00:3",
                      "secondary_exposures": [
                        {
                          "gate": "segment:chatgpt_voice_team",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:friends_and_family_nda",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:chatgpt_paid_users",
                          "gateValue": "true",
                          "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                        },
                        {
                          "gate": "chatgpt-web-voice-message-playback",
                          "gateValue": "true",
                          "ruleID": "6VUF6Z1JaUKZF7RS6uSjUu:100.00:6"
                        }
                      ]
                    },
                    "211801712": {
                      "name": "211801712",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "226799327": {
                      "name": "226799327",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "437245079": {
                      "name": "437245079",
                      "value": true,
                      "rule_id": "1KKPtXmOjxZamlgM2DzXlg",
                      "secondary_exposures": []
                    },
                    "458009956": {
                      "name": "458009956",
                      "value": true,
                      "rule_id": "6LgEwykI8hHnFAQ5EkKTX2:100.00:1",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_free_preview_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_paid_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "544659310": {
                      "name": "544659310",
                      "value": true,
                      "rule_id": "1Opz1br1XsEM6Uajp6F1XU:100.00:3",
                      "secondary_exposures": []
                    },
                    "607723475": {
                      "name": "607723475",
                      "value": false,
                      "rule_id": "2e0cON2LQDfCjkZQ6yzHg5",
                      "secondary_exposures": []
                    },
                    "716475478": {
                      "name": "716475478",
                      "value": true,
                      "rule_id": "48mToGJmOtQ5aRx9MCjOor",
                      "secondary_exposures": []
                    },
                    "716722001": {
                      "name": "716722001",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "730579107": {
                      "name": "730579107",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "740954505": {
                      "name": "740954505",
                      "value": true,
                      "rule_id": "2Ey97LJY7QymPfJfxULJwr:85.00:5",
                      "secondary_exposures": [
                        {
                          "gate": "segment:anonymous_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "internal-employee-only-chatgpt-team",
                          "gateValue": "false",
                          "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                        },
                        {
                          "gate": "segment:enterprise",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_free_preview_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_paid_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "771993664": {
                      "name": "771993664",
                      "value": false,
                      "rule_id": "1dDdxrXy1KadHtYusuGRgR:0.00:5",
                      "secondary_exposures": []
                    },
                    "857491970": {
                      "name": "857491970",
                      "value": true,
                      "rule_id": "2LpyNTIfbYC4R6XS6dfdIL",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "859327441": {
                      "name": "859327441",
                      "value": false,
                      "rule_id": "2kLMbBkcSZ8TkzHgzgbyKV",
                      "secondary_exposures": []
                    },
                    "874794369": {
                      "name": "874794369",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "1026575464": {
                      "name": "1026575464",
                      "value": true,
                      "rule_id": "1f4cSeenpcmWQ9eMKKQhF5",
                      "secondary_exposures": []
                    },
                    "1056591003": {
                      "name": "1056591003",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:internal_oai_gptv_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:chatgpt_unpaid_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-mainline-model-transition-unpaid-sahara-v",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "1070272635": {
                      "name": "1070272635",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:friends_and_family_nda",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:feather_memory_trainers",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "1123641887": {
                      "name": "1123641887",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt_anon_chat_enabled",
                          "gateValue": "true",
                          "ruleID": "22nVhoL17eyMvGWgFrDfZe"
                        }
                      ]
                    },
                    "1123642877": {
                      "name": "1123642877",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt_anon_chat_enabled",
                          "gateValue": "true",
                          "ruleID": "22nVhoL17eyMvGWgFrDfZe"
                        }
                      ]
                    },
                    "1162285750": {
                      "name": "1162285750",
                      "value": false,
                      "rule_id": "5SD2EhDYytn75oTAjOCSiL",
                      "secondary_exposures": []
                    },
                    "1198030896": {
                      "name": "1198030896",
                      "value": true,
                      "rule_id": "2aaGr86KHS6GUSs6I9F2cV",
                      "secondary_exposures": []
                    },
                    "1200224890": {
                      "name": "1200224890",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:anonymous_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:enterprise",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "internal-employee-only-chatgpt-team",
                          "gateValue": "false",
                          "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                        },
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_free_preview_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_paid_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt_fruit_juice",
                          "gateValue": "true",
                          "ruleID": "2Ey97LJY7QymPfJfxULJwr:85.00:5"
                        }
                      ]
                    },
                    "1306782371": {
                      "name": "1306782371",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "1426009137": {
                      "name": "1426009137",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:chatgpt_voice_team",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "1428536080": {
                      "name": "1428536080",
                      "value": false,
                      "rule_id": "2lFzVsatjvpp5kR8PhJaEl",
                      "secondary_exposures": []
                    },
                    "1568195116": {
                      "name": "1568195116",
                      "value": false,
                      "rule_id": "KKhehNjRxNHoIWKNkqpU2",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "1654633603": {
                      "name": "1654633603",
                      "value": true,
                      "rule_id": "2WWYf2mhNtJ9juDjcM3gmE",
                      "secondary_exposures": [
                        {
                          "gate": "segment:anonymous_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "internal-employee-only-chatgpt-team",
                          "gateValue": "false",
                          "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                        },
                        {
                          "gate": "segment:enterprise",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_free_preview_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_paid_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt_fruit_juice",
                          "gateValue": "true",
                          "ruleID": "2Ey97LJY7QymPfJfxULJwr:85.00:5"
                        }
                      ]
                    },
                    "1703773460": {
                      "name": "1703773460",
                      "value": false,
                      "rule_id": "63qNN5LSdeNXfSPo2Mwuv2:0.00:2",
                      "secondary_exposures": []
                    },
                    "1820151945": {
                      "name": "1820151945",
                      "value": true,
                      "rule_id": "3dG8z8CNFtgpXJNyOi8ijR",
                      "secondary_exposures": [
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:model_switcher_team",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:chatgpt_paid_users",
                          "gateValue": "true",
                          "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                        }
                      ]
                    },
                    "1825130190": {
                      "name": "1825130190",
                      "value": true,
                      "rule_id": "Nef2uMceNUF9U3ZYwSbpD",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "1828361965": {
                      "name": "1828361965",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "1852364556": {
                      "name": "1852364556",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "1880477506": {
                      "name": "1880477506",
                      "value": true,
                      "rule_id": "7fFsLeJu6pyCq9tRWmX9em:100.00:6",
                      "secondary_exposures": []
                    },
                    "1883608273": {
                      "name": "1883608273",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "1923022511": {
                      "name": "1923022511",
                      "value": true,
                      "rule_id": "6VUF6Z1JaUKZF7RS6uSjUu:100.00:6",
                      "secondary_exposures": [
                        {
                          "gate": "segment:chatgpt_voice_team",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:friends_and_family_nda",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:chatgpt_paid_users",
                          "gateValue": "true",
                          "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                        }
                      ]
                    },
                    "1938289220": {
                      "name": "1938289220",
                      "value": false,
                      "rule_id": "1vWEhLfJ39edDu8ZjE19dl:50.00:4",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:contractors",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:sidekick_windows_contractors",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:friends_and_family_nda",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:enterprise",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:chatgpt_paid_users",
                          "gateValue": "true",
                          "ruleID": "5qtkUmze51QU4V1SsDjPcb"
                        }
                      ]
                    },
                    "2000076788": {
                      "name": "2000076788",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:contractors",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2004295676": {
                      "name": "2004295676",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:sonic-testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-sonic-internal",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-search-demo-ui",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2046333385": {
                      "name": "2046333385",
                      "value": false,
                      "rule_id": "4nM2ehmgoDQIv69B0zohb6",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2166885617": {
                      "name": "2166885617",
                      "value": false,
                      "rule_id": "1ORnWBMVyZ42eAp03fe5Nr",
                      "secondary_exposures": []
                    },
                    "2173548801": {
                      "name": "2173548801",
                      "value": true,
                      "rule_id": "22nVhoL17eyMvGWgFrDfZe",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2438078654": {
                      "name": "2438078654",
                      "value": true,
                      "rule_id": "VNBdYfNSNNUyDdMpaiaZ9:100.00:3",
                      "secondary_exposures": []
                    },
                    "2449630478": {
                      "name": "2449630478",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "2542814321": {
                      "name": "2542814321",
                      "value": false,
                      "rule_id": "6Usmt3Th0mkUkVQJ0oyeST:0.00:1",
                      "secondary_exposures": []
                    },
                    "2546838099": {
                      "name": "2546838099",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:sonic-testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-sonic-internal",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2562876640": {
                      "name": "2562876640",
                      "value": true,
                      "rule_id": "326czTZeZ0RX0ypR0c5Bb6:100.00:15",
                      "secondary_exposures": []
                    },
                    "2689530298": {
                      "name": "2689530298",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2865702999": {
                      "name": "2865702999",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2865703032": {
                      "name": "2865703032",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2893926039": {
                      "name": "2893926039",
                      "value": true,
                      "rule_id": "1XOa3iabw20dgZn5lg5HTh",
                      "secondary_exposures": []
                    },
                    "2947776398": {
                      "name": "2947776398",
                      "value": false,
                      "rule_id": "2YzNwxxdROEWn1lPwJbYtr",
                      "secondary_exposures": []
                    },
                    "3011278998": {
                      "name": "3011278998",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "3060586957": {
                      "name": "3060586957",
                      "value": true,
                      "rule_id": "1UTG3o9yKWQjDvhcNL1pSX",
                      "secondary_exposures": []
                    },
                    "3134329717": {
                      "name": "3134329717",
                      "value": false,
                      "rule_id": "2UWRV6bObClvvj8jLkP7kJ",
                      "secondary_exposures": []
                    },
                    "3254335143": {
                      "name": "3254335143",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": []
                    },
                    "3320547351": {
                      "name": "3320547351",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "3376455464": {
                      "name": "3376455464",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "3455796462": {
                      "name": "3455796462",
                      "value": true,
                      "rule_id": "YbNn9kuhmMn3DdnzEDc6G",
                      "secondary_exposures": []
                    },
                    "3483957186": {
                      "name": "3483957186",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:anonymous_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "3501550327": {
                      "name": "3501550327",
                      "value": false,
                      "rule_id": "HEU9zVkuwUfmXHhGcckPQ",
                      "secondary_exposures": []
                    },
                    "3511639124": {
                      "name": "3511639124",
                      "value": true,
                      "rule_id": "1y9t5DBS0nz9mqbQRCrmCP:100.00:5",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:plus_qa_-_gizmos",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:store_dogfooding",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:anonymous_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:gizmo_store_holdout_exceptions",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "3643162575": {
                      "name": "3643162575",
                      "value": true,
                      "rule_id": "56j9Tu0CGzcoiVbnTxuR7O",
                      "secondary_exposures": []
                    },
                    "3700615661": {
                      "name": "3700615661",
                      "value": true,
                      "rule_id": "66covjaoZoe9pQR4I68jOB",
                      "secondary_exposures": []
                    },
                    "3912844760": {
                      "name": "3912844760",
                      "value": false,
                      "rule_id": "default",
                      "secondary_exposures": [
                        {
                          "gate": "segment:european_legal_restricted",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "3922145230": {
                      "name": "3922145230",
                      "value": false,
                      "rule_id": "14DZA2LumaPqAdCo52CrUB",
                      "secondary_exposures": [
                        {
                          "gate": "segment:internal_ips",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "3941367332": {
                      "name": "3941367332",
                      "value": true,
                      "rule_id": "3tesvDaP0kRy0JAXNdSJJn:100.00:3",
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                          "gateValue": "true",
                          "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                        }
                      ]
                    },
                    "4226692983": {
                      "name": "4226692983",
                      "value": true,
                      "rule_id": "6sEu91zwlBGSKOqFiNpGlA:100.00:2",
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:enterprise",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:teams_plan",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    }
                  },
                  "dynamic_configs": {
                    "75342665": {
                      "name": "75342665",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": ["is_anon_chat_enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "75342698": {
                      "name": "75342698",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": ["is_anon_chat_enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "101132775": {
                      "name": "101132775",
                      "value": { "show_new_ui": true },
                      "group": "launchedGroup",
                      "rule_id": "launchedGroup",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false
                    },
                    "171317775": {
                      "name": "171317775",
                      "value": {},
                      "group": "prestart",
                      "rule_id": "prestart",
                      "is_device_based": true,
                      "secondary_exposures": [],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false
                    },
                    "340466392": {
                      "name": "340466392",
                      "value": {
                        "google.com": [
                          "google.com",
                          "googleapis.com",
                          "www.googleapis.com",
                          "oauth2.googleapis.com",
                          "cloudfunctions.net"
                        ],
                        "microsoftonline.com": [
                          "microsoftonline.com",
                          "azurewebsites.net",
                          "graph.microsoft.com",
                          "microsoft.com",
                          "azure.com",
                          "sinequa.com"
                        ],
                        "adobelogin.com": ["adobe.com", "adobelogin.com"],
                        "useflorian.com": ["useflorian.com", "convex.site"],
                        "yahoo.com": ["yahooapis.com", "yahoo.com"],
                        "visualstudio.com": ["visualstudio.com", "azure.com"],
                        "atlassian.com": ["atlassian.net", "atlassian.com"]
                      },
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": true,
                      "secondary_exposures": []
                    },
                    "445876862": {
                      "name": "445876862",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [
                        "is_anon_chat_enabled",
                        "is_no_auth_welcome_modal_enabled",
                        "no_auth_soft_rate_limit",
                        "no_auth_hard_rate_limit",
                        "should_show_no_auth_signup_banner"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "855925714": {
                      "name": "855925714",
                      "value": {
                        "should_simplify_modal": false,
                        "is_simplified_sharing_modal_enabled": false,
                        "is_social_share_options_enabled": false,
                        "is_update_shared_links_enabled": false,
                        "is_discoverability_toggle_enabled": false,
                        "show_copylink_state_if_no_updates": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_simplified_sharing_modal_enabled",
                        "is_social_share_options_enabled",
                        "is_update_shared_links_enabled",
                        "is_discoverability_toggle_enabled"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "889327355": {
                      "name": "889327355",
                      "value": { "is-enabled": true },
                      "group": "layerAssignment",
                      "rule_id": "layerAssignment",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": ["is-enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": true,
                      "is_in_layer": true
                    },
                    "1193714591": {
                      "name": "1193714591",
                      "value": {
                        "is_dynamic_model_enabled": false,
                        "show_message_model_info": false,
                        "show_message_regenerate_model_selector": true,
                        "is_conversation_model_switching_allowed": true,
                        "show_rate_limit_downgrade_banner": true,
                        "config": {
                          "impl": "rm_score",
                          "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                          "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                        },
                        "show_message_regenerate_model_selector_on_every_message": true,
                        "is_AG8PqS2q_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_dynamic_model_enabled",
                        "show_message_regenerate_model_selector",
                        "is_conversation_model_switching_allowed",
                        "show_rate_limit_downgrade_banner"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "1193714592": {
                      "name": "1193714592",
                      "value": {
                        "is_dynamic_model_enabled": false,
                        "show_message_model_info": false,
                        "show_message_regenerate_model_selector": true,
                        "is_conversation_model_switching_allowed": true,
                        "show_rate_limit_downgrade_banner": true,
                        "config": {
                          "impl": "rm_score",
                          "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                          "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                        },
                        "show_message_regenerate_model_selector_on_every_message": true,
                        "is_AG8PqS2q_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_dynamic_model_enabled",
                        "show_message_regenerate_model_selector",
                        "is_conversation_model_switching_allowed",
                        "show_rate_limit_downgrade_banner",
                        "config"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "1193714593": {
                      "name": "1193714593",
                      "value": {
                        "is_dynamic_model_enabled": false,
                        "show_message_model_info": false,
                        "show_message_regenerate_model_selector": true,
                        "is_conversation_model_switching_allowed": true,
                        "show_rate_limit_downgrade_banner": true,
                        "config": {
                          "impl": "rm_score",
                          "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                          "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                        },
                        "show_message_regenerate_model_selector_on_every_message": true,
                        "is_AG8PqS2q_enabled": false
                      },
                      "group": "prestart",
                      "rule_id": "prestart",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_dynamic_model_enabled",
                        "show_message_regenerate_model_selector",
                        "is_conversation_model_switching_allowed",
                        "show_rate_limit_downgrade_banner",
                        "config",
                        "show_message_regenerate_model_selector_on_every_message"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "1237904433": {
                      "name": "1237904433",
                      "value": { "is-enabled": true },
                      "group": "launchedGroup",
                      "rule_id": "launchedGroup",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": ["is-enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "1277879515": {
                      "name": "1277879515",
                      "value": { "show_ui": true },
                      "group": "launchedGroup",
                      "rule_id": "launchedGroup",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "segment:friends_and_family_nda",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:feather_memory_trainers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-allow-memory-use-admin",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false
                    },
                    "1312738797": {
                      "name": "1312738797",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "launchedGroup",
                      "rule_id": "launchedGroup",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": ["anon-is-spanish-translation-enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "1476759596": {
                      "name": "1476759596",
                      "value": {},
                      "group": "layerAssignment",
                      "rule_id": "layerAssignment",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "is_user_in_experiment": false,
                      "is_experiment_active": true
                    },
                    "1547924522": {
                      "name": "1547924522",
                      "value": { "show_ui": true, "enabled": true },
                      "group": "launchedGroup",
                      "rule_id": "launchedGroup",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt-paid-users",
                          "gateValue": "true",
                          "ruleID": "7stvavArTuHK8XvexxjCiu"
                        }
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false
                    },
                    "1887692490": {
                      "name": "1887692490",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [
                        "is_anon_chat_enabled",
                        "is_try_it_first_on_login_page_enabled"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "1908072088": {
                      "name": "1908072088",
                      "value": {
                        "blocking_modal": true,
                        "single_button_variant": true,
                        "dropdown_tooltip": false,
                        "conversation_model_switcher_badge": false
                      },
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": false,
                      "secondary_exposures": []
                    },
                    "2112835930": {
                      "name": "2112835930",
                      "value": {},
                      "group": "targetingGate",
                      "rule_id": "targetingGate",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "segment:anonymous_users",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:enterprise",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "internal-employee-only-chatgpt-team",
                          "gateValue": "false",
                          "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                        },
                        {
                          "gate": "segment:plusqa_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_free_preview_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:scallion_paid_testers",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt_fruit_juice",
                          "gateValue": "true",
                          "ruleID": "2Ey97LJY7QymPfJfxULJwr:85.00:5"
                        },
                        {
                          "gate": "chatgpt_fruit_juice_inverse",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": true
                    },
                    "2144828235": {
                      "name": "2144828235",
                      "value": {},
                      "group": "layerAssignment",
                      "rule_id": "layerAssignment",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "is_user_in_experiment": false,
                      "is_experiment_active": true
                    },
                    "2176179500": {
                      "name": "2176179500",
                      "value": { "should_enable_zh_tw": false },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": ["should_enable_zh_tw"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2331692225": {
                      "name": "2331692225",
                      "value": {},
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-browse-sahara",
                          "gateValue": "false",
                          "ruleID": "default"
                        }
                      ]
                    },
                    "2370410632": {
                      "name": "2370410632",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": ["is_anon_chat_enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2370410634": {
                      "name": "2370410634",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": ["is_anon_chat_enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2575540972": {
                      "name": "2575540972",
                      "value": {
                        "is_dynamic_model_enabled": false,
                        "show_message_model_info": false,
                        "show_message_regenerate_model_selector": true,
                        "is_conversation_model_switching_allowed": true,
                        "show_rate_limit_downgrade_banner": true,
                        "config": {
                          "impl": "rm_score",
                          "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                          "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                        },
                        "show_message_regenerate_model_selector_on_every_message": true,
                        "is_AG8PqS2q_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_dynamic_model_enabled",
                        "show_message_regenerate_model_selector",
                        "is_conversation_model_switching_allowed",
                        "show_rate_limit_downgrade_banner"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2576123431": {
                      "name": "2576123431",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [
                        "is_anon_chat_enabled",
                        "is_anon_chat_enabled_for_new_users_only"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2635734272": {
                      "name": "2635734272",
                      "value": {
                        "should_simplify_modal": false,
                        "is_simplified_sharing_modal_enabled": false,
                        "is_social_share_options_enabled": false,
                        "is_update_shared_links_enabled": false,
                        "is_discoverability_toggle_enabled": false,
                        "show_copylink_state_if_no_updates": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_simplified_sharing_modal_enabled",
                        "is_social_share_options_enabled",
                        "is_update_shared_links_enabled",
                        "is_discoverability_toggle_enabled"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2745767556": {
                      "name": "2745767556",
                      "value": {
                        "is_team_enabled": true,
                        "is_yearly_plus_subscription_enabled": false,
                        "is_split_between_personal_and_business_enabled": false
                      },
                      "group": "prestart",
                      "rule_id": "prestart",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        }
                      ],
                      "explicit_parameters": [
                        "is_split_between_personal_and_business_enabled"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2811617088": {
                      "name": "2811617088",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [
                        "is_anon_chat_enabled",
                        "is_no_auth_welcome_modal_enabled",
                        "no_auth_soft_rate_limit",
                        "no_auth_hard_rate_limit",
                        "should_show_no_auth_signup_banner"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "2891383514": {
                      "name": "2891383514",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "layerAssignment",
                      "rule_id": "layerAssignment",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [
                        "no_auth_soft_rate_limit",
                        "is_no_auth_soft_rate_limit_modal_enabled"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": true,
                      "is_in_layer": true
                    },
                    "2919366999": {
                      "name": "2919366999",
                      "value": {
                        "should_simplify_modal": false,
                        "is_simplified_sharing_modal_enabled": false,
                        "is_social_share_options_enabled": false,
                        "is_update_shared_links_enabled": false,
                        "is_discoverability_toggle_enabled": false,
                        "show_copylink_state_if_no_updates": false
                      },
                      "group": "3fdACLHPKSOwzpRi35OaGG",
                      "rule_id": "3fdACLHPKSOwzpRi35OaGG",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                          "gateValue": "true",
                          "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                        }
                      ],
                      "explicit_parameters": ["show_copylink_state_if_no_updates"],
                      "is_user_in_experiment": true,
                      "is_experiment_active": true,
                      "is_in_layer": true
                    },
                    "2980740905": {
                      "name": "2980740905",
                      "value": {
                        "is_team_enabled": true,
                        "is_yearly_plus_subscription_enabled": false,
                        "is_split_between_personal_and_business_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        }
                      ],
                      "explicit_parameters": ["is_yearly_plus_subscription_enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "3168246095": {
                      "name": "3168246095",
                      "value": { "gizmo_ids": [] },
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": false,
                      "secondary_exposures": []
                    },
                    "3299284219": {
                      "name": "3299284219",
                      "value": {
                        "is_dynamic_model_enabled": false,
                        "show_message_model_info": false,
                        "show_message_regenerate_model_selector": true,
                        "is_conversation_model_switching_allowed": true,
                        "show_rate_limit_downgrade_banner": true,
                        "config": {
                          "impl": "rm_score",
                          "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                          "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                        },
                        "show_message_regenerate_model_selector_on_every_message": true,
                        "is_AG8PqS2q_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_dynamic_model_enabled",
                        "show_message_regenerate_model_selector",
                        "is_conversation_model_switching_allowed",
                        "show_rate_limit_downgrade_banner",
                        "show_message_regenerate_model_selector_on_every_message"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "3340924940": {
                      "name": "3340924940",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [
                        "is_anon_chat_enabled",
                        "should_show_anon_login_header_on_desktop"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "3414158273": {
                      "name": "3414158273",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "launchedGroup",
                      "rule_id": "launchedGroup",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [
                        "is_try_it_first_on_login_page_enabled",
                        "is_no_auth_welcome_back_modal_enabled"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "3414159263": {
                      "name": "3414159263",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "layerAssignment",
                      "rule_id": "layerAssignment",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": ["is_no_auth_welcome_back_modal_enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": true,
                      "is_in_layer": true
                    },
                    "3758098899": {
                      "name": "3758098899",
                      "value": {
                        "is_team_enabled": true,
                        "is_yearly_plus_subscription_enabled": false,
                        "is_split_between_personal_and_business_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        }
                      ],
                      "explicit_parameters": ["is_team_enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "3758099832": {
                      "name": "3758099832",
                      "value": {},
                      "group": "prestart",
                      "rule_id": "prestart",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false
                    },
                    "3917128856": {
                      "name": "3917128856",
                      "value": {
                        "is_dynamic_model_enabled": false,
                        "show_message_model_info": false,
                        "show_message_regenerate_model_selector": true,
                        "is_conversation_model_switching_allowed": true,
                        "show_rate_limit_downgrade_banner": true,
                        "config": {
                          "impl": "rm_score",
                          "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                          "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                        },
                        "show_message_regenerate_model_selector_on_every_message": true,
                        "is_AG8PqS2q_enabled": false
                      },
                      "group": "abandoned",
                      "rule_id": "abandoned",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [
                        "is_dynamic_model_enabled",
                        "show_message_regenerate_model_selector",
                        "is_conversation_model_switching_allowed",
                        "show_rate_limit_downgrade_banner",
                        "config"
                      ],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    },
                    "4204547666": {
                      "name": "4204547666",
                      "value": { "is-enabled": true },
                      "group": "Nzc6Xnht6tIVmb48Ejg1T:override",
                      "rule_id": "Nzc6Xnht6tIVmb48Ejg1T:override",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "segment:internal_ips",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-localization-preview",
                          "gateValue": "false",
                          "ruleID": "14DZA2LumaPqAdCo52CrUB"
                        },
                        {
                          "gate": "chatgpt-localization-allowlist",
                          "gateValue": "true",
                          "ruleID": "66covjaoZoe9pQR4I68jOB"
                        }
                      ],
                      "explicit_parameters": ["is-enabled"],
                      "is_user_in_experiment": false,
                      "is_experiment_active": false,
                      "is_in_layer": true
                    }
                  },
                  "layer_configs": {
                    "468168202": {
                      "name": "468168202",
                      "value": {
                        "is_team_enabled": true,
                        "is_yearly_plus_subscription_enabled": false,
                        "is_split_between_personal_and_business_enabled": false
                      },
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        }
                      ],
                      "explicit_parameters": [],
                      "undelegated_secondary_exposures": [
                        {
                          "gate": "chatgpt_yearly_plus_subscription_holdout_240117",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        }
                      ]
                    },
                    "1238742812": {
                      "name": "1238742812",
                      "value": { "should_enable_zh_tw": false },
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [],
                      "undelegated_secondary_exposures": []
                    },
                    "1547743984": {
                      "name": "1547743984",
                      "value": {
                        "should_simplify_modal": false,
                        "is_simplified_sharing_modal_enabled": false,
                        "is_social_share_options_enabled": false,
                        "is_update_shared_links_enabled": false,
                        "is_discoverability_toggle_enabled": false,
                        "show_copylink_state_if_no_updates": false
                      },
                      "group": "3fdACLHPKSOwzpRi35OaGG",
                      "rule_id": "3fdACLHPKSOwzpRi35OaGG",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                          "gateValue": "true",
                          "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                        }
                      ],
                      "explicit_parameters": ["show_copylink_state_if_no_updates"],
                      "allocated_experiment_name": "2919366999",
                      "is_experiment_active": true,
                      "is_user_in_experiment": true,
                      "undelegated_secondary_exposures": []
                    },
                    "2723963139": {
                      "name": "2723963139",
                      "value": {
                        "is_dynamic_model_enabled": false,
                        "show_message_model_info": false,
                        "show_message_regenerate_model_selector": true,
                        "is_conversation_model_switching_allowed": true,
                        "show_rate_limit_downgrade_banner": true,
                        "config": {
                          "impl": "rm_score",
                          "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                          "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                        },
                        "show_message_regenerate_model_selector_on_every_message": true,
                        "is_AG8PqS2q_enabled": false
                      },
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": false,
                      "secondary_exposures": [],
                      "explicit_parameters": [],
                      "undelegated_secondary_exposures": []
                    },
                    "3048336830": {
                      "name": "3048336830",
                      "value": { "is-enabled": true },
                      "group": "Nzc6Xnht6tIVmb48Ejg1T:override",
                      "rule_id": "Nzc6Xnht6tIVmb48Ejg1T:override",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "segment:internal_ips",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-localization-preview",
                          "gateValue": "false",
                          "ruleID": "14DZA2LumaPqAdCo52CrUB"
                        },
                        {
                          "gate": "chatgpt-localization-allowlist",
                          "gateValue": "true",
                          "ruleID": "66covjaoZoe9pQR4I68jOB"
                        }
                      ],
                      "explicit_parameters": ["is-enabled"],
                      "allocated_experiment_name": "4204547666",
                      "is_experiment_active": false,
                      "is_user_in_experiment": false,
                      "undelegated_secondary_exposures": [
                        {
                          "gate": "segment:internal_ips",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "segment:employees",
                          "gateValue": "false",
                          "ruleID": "default"
                        },
                        {
                          "gate": "chatgpt-localization-preview",
                          "gateValue": "false",
                          "ruleID": "14DZA2LumaPqAdCo52CrUB"
                        },
                        {
                          "gate": "chatgpt-localization-allowlist",
                          "gateValue": "true",
                          "ruleID": "66covjaoZoe9pQR4I68jOB"
                        }
                      ]
                    },
                    "3637408529": {
                      "name": "3637408529",
                      "value": {
                        "is_anon_chat_enabled": true,
                        "anon_composer_display_variant": "default",
                        "anon-is-spanish-translation-enabled": true,
                        "should_show_anon_login_header_on_desktop": false,
                        "is_anon_chat_enabled_for_new_users_only": false,
                        "is_try_it_first_on_login_page_enabled": false,
                        "is_no_auth_welcome_modal_enabled": false,
                        "no_auth_soft_rate_limit": 1200,
                        "no_auth_hard_rate_limit": 1200,
                        "should_show_no_auth_signup_banner": true,
                        "is_no_auth_welcome_back_modal_enabled": true,
                        "is_no_auth_soft_rate_limit_modal_enabled": false
                      },
                      "group": "default",
                      "rule_id": "default",
                      "is_device_based": false,
                      "secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ],
                      "explicit_parameters": [],
                      "undelegated_secondary_exposures": [
                        {
                          "gate": "chatgpt_anon_chat_holdout_20240227",
                          "gateValue": "false",
                          "ruleID": "disabled"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_gate_20240328",
                          "gateValue": "true",
                          "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                        },
                        {
                          "gate": "chatgpt_no_auth_holdout_20240328",
                          "gateValue": "false",
                          "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                        }
                      ]
                    }
                  },
                  "sdkParams": {},
                  "has_updates": true,
                  "generator": "statsig-node-sdk",
                  "time": 0,
                  "evaluated_keys": {
                    "userID": "user-ieBaO6Zyq6qmnqo6k4qEyj9Q",
                    "customIDs": {
                      "WebAnonymousCookieID": "d6b6c8fe-69ff-4722-9124-15d50e750b62",
                      "DeviceId": "d6b6c8fe-69ff-4722-9124-15d50e750b62",
                      "stableID": "d6b6c8fe-69ff-4722-9124-15d50e750b62"
                    }
                  },
                  "hash_used": "djb2",
                  "user_hash": "633226185"
                },
                "user": {
                  "country": "US",
                  "custom": {
                    "has_logged_in_before": true,
                    "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
                    "is_paid": true,
                    "auth_status": "logged_in",
                    "workspace_id": "53ebfd0f-3d00-4441-ab26-f45a6fd8f2d9"
                  },
                  "customIDs": {
                    "WebAnonymousCookieID": "d6b6c8fe-69ff-4722-9124-15d50e750b62",
                    "DeviceId": "d6b6c8fe-69ff-4722-9124-15d50e750b62",
                    "stableID": "d6b6c8fe-69ff-4722-9124-15d50e750b62"
                  },
                  "ip": "1.1.1.1",
                  "locale": "zh-CN",
                  "privateAttributes": { "email": "xyhelper@closeai.biz" },
                  "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
                  "userID": "user-ieBaO6Zyq6qmnqo6k4qEyj9Q"
                }
              },
              "authStatus": "logged_in",
              "serverPrimedAllowBrowserStorageValue": true,
              "showCookieConsentBanner": false,
              "isStorageComplianceEnabled": false,
              "ageVerificationDeadline": null,
              "ageVerificationType": "none",
              "logOutUser": false
            },
            "__N_SSP": true
          },
          "page": "/g/[gizmoId]",
          "query": { "gizmoId": "g-pmuQfob8d-image-generator" },
          "buildId": "-wRE4Obkm_QOW7PLn1x21",
          "assetPrefix": "https://cdn.oaistatic.com",
          "isFallback": false,
          "isExperimentalCompile": false,
          "gssp": true,
          "scriptLoader": []
        }
        `
)