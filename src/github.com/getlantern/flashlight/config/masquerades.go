package config

import "github.com/getlantern/fronted"

var defaultTrustedCAs = []*CA{
	&CA{
		CommonName: "GlobalSign Root CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG\nA1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv\nb3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw\nMDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i\nYWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT\naWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ\njc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp\nxy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp\n1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG\nsnUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ\nU26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8\n9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E\nBTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B\nAQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz\nyj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE\n38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP\nAbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad\nDKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME\nHMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "AddTrust External CA Root",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIENjCCAx6gAwIBAgIBATANBgkqhkiG9w0BAQUFADBvMQswCQYDVQQGEwJTRTEU\nMBIGA1UEChMLQWRkVHJ1c3QgQUIxJjAkBgNVBAsTHUFkZFRydXN0IEV4dGVybmFs\nIFRUUCBOZXR3b3JrMSIwIAYDVQQDExlBZGRUcnVzdCBFeHRlcm5hbCBDQSBSb290\nMB4XDTAwMDUzMDEwNDgzOFoXDTIwMDUzMDEwNDgzOFowbzELMAkGA1UEBhMCU0Ux\nFDASBgNVBAoTC0FkZFRydXN0IEFCMSYwJAYDVQQLEx1BZGRUcnVzdCBFeHRlcm5h\nbCBUVFAgTmV0d29yazEiMCAGA1UEAxMZQWRkVHJ1c3QgRXh0ZXJuYWwgQ0EgUm9v\ndDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALf3GjPm8gAELTngTlvt\nH7xsD821+iO2zt6bETOXpClMfZOfvUq8k+0DGuOPz+VtUFrWlymUWoCwSXrbLpX9\nuMq/NzgtHj6RQa1wVsfwTz/oMp50ysiQVOnGXw94nZpAPA6sYapeFI+eh6FqUNzX\nmk6vBbOmcZSccbNQYArHE504B4YCqOmoaSYYkKtMsE8jqzpPhNjfzp/haW+710LX\na0Tkx63ubUFfclpxCDezeWWkWaCUN/cALw3CknLa0Dhy2xSoRcRdKn23tNbE7qzN\nE0S3ySvdQwAl+mG5aWpYIxG3pzOPVnVZ9c0p10a3CitlttNCbxWyuHv77+ldU9U0\nWicCAwEAAaOB3DCB2TAdBgNVHQ4EFgQUrb2YejS0Jvf6xCZU7wO94CTLVBowCwYD\nVR0PBAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wgZkGA1UdIwSBkTCBjoAUrb2YejS0\nJvf6xCZU7wO94CTLVBqhc6RxMG8xCzAJBgNVBAYTAlNFMRQwEgYDVQQKEwtBZGRU\ncnVzdCBBQjEmMCQGA1UECxMdQWRkVHJ1c3QgRXh0ZXJuYWwgVFRQIE5ldHdvcmsx\nIjAgBgNVBAMTGUFkZFRydXN0IEV4dGVybmFsIENBIFJvb3SCAQEwDQYJKoZIhvcN\nAQEFBQADggEBALCb4IUlwtYj4g+WBpKdQZic2YR5gdkeWxQHIzZlj7DYd7usQWxH\nYINRsPkyPef89iYTx4AWpb9a/IfPeHmJIZriTAcKhjW88t5RxNKWt9x+Tu5w/Rw5\n6wwCURQtjr0W4MHfRnXnJK3s9EK0hZNwEGe6nQY1ShjTK3rMUUKhemPR5ruhxSvC\nNr4TDea9Y355e6cJDUCrat2PisP29owaQgVR1EX1n6diIWgVIEM8med8vSTYqZEX\nc4g/VhsxOBi0cQ+azcgOno4uG+GMmIPLHzHxREzGBHNJdmAPx/i9F4BrLunMTA5a\nmnkPIAou1Z5jJh5VkpTYghdae9C8x49OhgQ=\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "Go Daddy Root Certificate Authority - G2",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDxTCCAq2gAwIBAgIBADANBgkqhkiG9w0BAQsFADCBgzELMAkGA1UEBhMCVVMx\nEDAOBgNVBAgTB0FyaXpvbmExEzARBgNVBAcTClNjb3R0c2RhbGUxGjAYBgNVBAoT\nEUdvRGFkZHkuY29tLCBJbmMuMTEwLwYDVQQDEyhHbyBEYWRkeSBSb290IENlcnRp\nZmljYXRlIEF1dGhvcml0eSAtIEcyMB4XDTA5MDkwMTAwMDAwMFoXDTM3MTIzMTIz\nNTk1OVowgYMxCzAJBgNVBAYTAlVTMRAwDgYDVQQIEwdBcml6b25hMRMwEQYDVQQH\nEwpTY290dHNkYWxlMRowGAYDVQQKExFHb0RhZGR5LmNvbSwgSW5jLjExMC8GA1UE\nAxMoR28gRGFkZHkgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkgLSBHMjCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL9xYgjx+lk09xvJGKP3gElY6SKD\nE6bFIEMBO4Tx5oVJnyfq9oQbTqC023CYxzIBsQU+B07u9PpPL1kwIuerGVZr4oAH\n/PMWdYA5UXvl+TW2dE6pjYIT5LY/qQOD+qK+ihVqf94Lw7YZFAXK6sOoBJQ7Rnwy\nDfMAZiLIjWltNowRGLfTshxgtDj6AozO091GB94KPutdfMh8+7ArU6SSYmlRJQVh\nGkSBjCypQ5Yj36w6gZoOKcUcqeldHraenjAKOc7xiID7S13MMuyFYkMlNAJWJwGR\ntDtwKj9useiciAF9n9T521NtYJ2/LOdYq7hfRvzOxBsDPAnrSTFcaUaz4EcCAwEA\nAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYE\nFDqahQcQZyi27/a9BUFuIMGU2g/eMA0GCSqGSIb3DQEBCwUAA4IBAQCZ21151fmX\nWWcDYfF+OwYxdS2hII5PZYe096acvNjpL9DbWu7PdIxztDhC2gV7+AJ1uP2lsdeu\n9tfeE8tTEH6KRtGX+rcuKxGrkLAngPnon1rpN5+r5N9ss4UXnT3ZJE95kTXWXwTr\ngIOrmgIttRD02JDHBHNA7XIloKmf7J6raBKZV8aPEjoJpL1E/QYVN8Gb5DKj7Tjo\n2GTzLH4U/ALqn83/B2gX2yKQOC16jdFU8WnjXzPKej17CuPKf1855eJ1usV2GDPO\nLPAvTK33sefOT6jEm0pUBsV/fdUID+Ic/n4XuKxe9tQWskMJDE32p2u0mYRlynqI\n4uJEvlz36hz1\n-----END CERTIFICATE-----\n",
	},
	&CA{
		CommonName: "GeoTrust Global CA",
		Cert:       "-----BEGIN CERTIFICATE-----\nMIIDVDCCAjygAwIBAgIDAjRWMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT\nMRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i\nYWwgQ0EwHhcNMDIwNTIxMDQwMDAwWhcNMjIwNTIxMDQwMDAwWjBCMQswCQYDVQQG\nEwJVUzEWMBQGA1UEChMNR2VvVHJ1c3QgSW5jLjEbMBkGA1UEAxMSR2VvVHJ1c3Qg\nR2xvYmFsIENBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2swYYzD9\n9BcjGlZ+W988bDjkcbd4kdS8odhM+KhDtgPpTSEHCIjaWC9mOSm9BXiLnTjoBbdq\nfnGk5sRgprDvgOSJKA+eJdbtg/OtppHHmMlCGDUUna2YRpIuT8rxh0PBFpVXLVDv\niS2Aelet8u5fa9IAjbkU+BQVNdnARqN7csiRv8lVK83Qlz6cJmTM386DGXHKTubU\n1XupGc1V3sjs0l44U+VcT4wt/lAjNvxm5suOpDkZALeVAjmRCw7+OC7RHQWa9k0+\nbw8HHa8sHo9gOeL6NlMTOdReJivbPagUvTLrGAMoUgRx5aszPeE4uwc2hGKceeoW\nMPRfwCvocWvk+QIDAQABo1MwUTAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTA\nephojYn7qwVkDBF9qn1luMrMTjAfBgNVHSMEGDAWgBTAephojYn7qwVkDBF9qn1l\nuMrMTjANBgkqhkiG9w0BAQUFAAOCAQEANeMpauUvXVSOKVCUn5kaFOSPeCpilKIn\nZ57QzxpeR+nBsqTP3UEaBU6bS+5Kb1VSsyShNwrrZHYqLizz/Tt1kL/6cdjHPTfS\ntQWVYrmm3ok9Nns4d0iXrKYgjy6myQzCsplFAMfOEVEiIuCl6rYVSAlk6l5PdPcF\nPseKUgzbFbS9bZvlxrFUaKnjaZC2mqUPuLk/IH2uSrW4nOQdtqvmlKXBx4Ot2/Un\nhw4EbNX/3aBd7YdStysVAq45pmp06drE57xNNB6pXE0zX5IJL4hmXXeXxx12E6nV\n5fEWCRE11azbJHFwLJhWC9kXtNHjUStedejV0NxPNO3CBWaAocvmMw==\n-----END CERTIFICATE-----\n",
	},
}

var cloudflareMasquerades = []*fronted.Masquerade{
	&fronted.Masquerade{
		Domain:    "1news.az",
		IpAddress: "162.159.241.30",
	},
	&fronted.Masquerade{
		Domain:    "2ch.hk",
		IpAddress: "162.159.254.5",
	},
	&fronted.Masquerade{
		Domain:    "a2hosting.com",
		IpAddress: "198.41.189.199",
	},
	&fronted.Masquerade{
		Domain:    "abs-cbnnews.com",
		IpAddress: "104.16.26.177",
	},
	&fronted.Masquerade{
		Domain:    "addmefast.com",
		IpAddress: "198.41.190.157",
	},
	&fronted.Masquerade{
		Domain:    "adf.ly",
		IpAddress: "104.20.0.4",
	},
	&fronted.Masquerade{
		Domain:    "adfoc.us",
		IpAddress: "162.159.254.16",
	},
	&fronted.Masquerade{
		Domain:    "adlure.net",
		IpAddress: "190.93.241.94",
	},
	&fronted.Masquerade{
		Domain:    "affiliatetechnology.com",
		IpAddress: "198.41.190.51",
	},
	&fronted.Masquerade{
		Domain:    "al-akhbar.com",
		IpAddress: "162.159.244.97",
	},
	&fronted.Masquerade{
		Domain:    "alexaboostup.com",
		IpAddress: "198.41.249.254",
	},
	&fronted.Masquerade{
		Domain:    "almasryalyoum.com",
		IpAddress: "141.101.112.103",
	},
	&fronted.Masquerade{
		Domain:    "alrakoba.net",
		IpAddress: "198.41.184.73",
	},
	&fronted.Masquerade{
		Domain:    "alsumaria.tv",
		IpAddress: "190.93.243.50",
	},
	&fronted.Masquerade{
		Domain:    "alwatanvoice.com",
		IpAddress: "162.159.255.142",
	},
	&fronted.Masquerade{
		Domain:    "amazinglytimedphotos.com",
		IpAddress: "198.41.185.180",
	},
	&fronted.Masquerade{
		Domain:    "amino.dk",
		IpAddress: "198.41.189.121",
	},
	&fronted.Masquerade{
		Domain:    "anazahra.com",
		IpAddress: "162.159.253.7",
	},
	&fronted.Masquerade{
		Domain:    "any.gs",
		IpAddress: "162.159.240.58",
	},
	&fronted.Masquerade{
		Domain:    "appstorm.net",
		IpAddress: "162.159.244.165",
	},
	&fronted.Masquerade{
		Domain:    "aqarcity.com",
		IpAddress: "198.41.187.74",
	},
	&fronted.Masquerade{
		Domain:    "aqarmap.com",
		IpAddress: "162.159.249.95",
	},
	&fronted.Masquerade{
		Domain:    "arabseed.com",
		IpAddress: "198.41.186.132",
	},
	&fronted.Masquerade{
		Domain:    "arageek.com",
		IpAddress: "198.41.205.85",
	},
	&fronted.Masquerade{
		Domain:    "armorgames.com",
		IpAddress: "104.20.5.17",
	},
	&fronted.Masquerade{
		Domain:    "authorstream.com",
		IpAddress: "190.93.246.194",
	},
	&fronted.Masquerade{
		Domain:    "avpixlat.info",
		IpAddress: "190.93.243.137",
	},
	&fronted.Masquerade{
		Domain:    "babyou.com",
		IpAddress: "108.162.206.67",
	},
	&fronted.Masquerade{
		Domain:    "banahosting.com",
		IpAddress: "162.159.245.11",
	},
	&fronted.Masquerade{
		Domain:    "baykoreans.net",
		IpAddress: "190.93.242.11",
	},
	&fronted.Masquerade{
		Domain:    "bezuzyteczna.pl",
		IpAddress: "198.41.179.171",
	},
	&fronted.Masquerade{
		Domain:    "bikroy.com",
		IpAddress: "104.16.23.214",
	},
	&fronted.Masquerade{
		Domain:    "bitcoinzebra.com",
		IpAddress: "104.20.6.88",
	},
	&fronted.Masquerade{
		Domain:    "bittrex.com",
		IpAddress: "162.159.246.225",
	},
	&fronted.Masquerade{
		Domain:    "bizimyol.info",
		IpAddress: "190.93.240.19",
	},
	&fronted.Masquerade{
		Domain:    "blabbermouth.net",
		IpAddress: "162.159.246.184",
	},
	&fronted.Masquerade{
		Domain:    "bleepingcomputer.com",
		IpAddress: "190.93.240.117",
	},
	&fronted.Masquerade{
		Domain:    "brainstorm9.com.br",
		IpAddress: "162.159.251.96",
	},
	&fronted.Masquerade{
		Domain:    "btc-e.com",
		IpAddress: "141.101.121.194",
	},
	&fronted.Masquerade{
		Domain:    "bubblews.com",
		IpAddress: "190.93.240.103",
	},
	&fronted.Masquerade{
		Domain:    "bukkit.org",
		IpAddress: "190.93.244.100",
	},
	&fronted.Masquerade{
		Domain:    "businessinsider.com.au",
		IpAddress: "190.93.247.134",
	},
	&fronted.Masquerade{
		Domain:    "cairokora.com",
		IpAddress: "104.16.3.117",
	},
	&fronted.Masquerade{
		Domain:    "careers360.com",
		IpAddress: "162.159.243.132",
	},
	&fronted.Masquerade{
		Domain:    "catracalivre.com.br",
		IpAddress: "198.41.247.124",
	},
	&fronted.Masquerade{
		Domain:    "censor.net.ua",
		IpAddress: "198.41.191.113",
	},
	&fronted.Masquerade{
		Domain:    "chinabuye.com",
		IpAddress: "198.41.186.203",
	},
	&fronted.Masquerade{
		Domain:    "cihan.com.tr",
		IpAddress: "104.16.2.7",
	},
	&fronted.Masquerade{
		Domain:    "citypatras.gr",
		IpAddress: "162.159.244.227",
	},
	&fronted.Masquerade{
		Domain:    "cleanfiles.net",
		IpAddress: "141.101.112.47",
	},
	&fronted.Masquerade{
		Domain:    "clixsense.com",
		IpAddress: "198.41.188.40",
	},
	&fronted.Masquerade{
		Domain:    "cloudify.cc",
		IpAddress: "162.159.252.62",
	},
	&fronted.Masquerade{
		Domain:    "coinmarketcap.com",
		IpAddress: "198.41.249.182",
	},
	&fronted.Masquerade{
		Domain:    "col3negoriginal.lk",
		IpAddress: "141.101.123.10",
	},
	&fronted.Masquerade{
		Domain:    "collective-evolution.com",
		IpAddress: "198.41.190.248",
	},
	&fronted.Masquerade{
		Domain:    "conversionxl.com",
		IpAddress: "162.159.242.170",
	},
	&fronted.Masquerade{
		Domain:    "convinceandconvert.com",
		IpAddress: "141.101.124.136",
	},
	&fronted.Masquerade{
		Domain:    "cpagrip.com",
		IpAddress: "198.41.184.139",
	},
	&fronted.Masquerade{
		Domain:    "cpasbien.pe",
		IpAddress: "104.16.17.124",
	},
	&fronted.Masquerade{
		Domain:    "cssmenumaker.com",
		IpAddress: "162.159.250.136",
	},
	&fronted.Masquerade{
		Domain:    "cuevana2.tv",
		IpAddress: "162.159.242.105",
	},
	&fronted.Masquerade{
		Domain:    "curse.com",
		IpAddress: "190.93.246.101",
	},
	&fronted.Masquerade{
		Domain:    "cursecdn.com",
		IpAddress: "198.41.208.103",
	},
	&fronted.Masquerade{
		Domain:    "dangerousminds.net",
		IpAddress: "108.162.204.89",
	},
	&fronted.Masquerade{
		Domain:    "dealcatcher.com",
		IpAddress: "162.159.249.16",
	},
	&fronted.Masquerade{
		Domain:    "delivery-club.ru",
		IpAddress: "185.17.186.230",
	},
	&fronted.Masquerade{
		Domain:    "demotywatory.pl",
		IpAddress: "95.211.149.169",
	},
	&fronted.Masquerade{
		Domain:    "deutsche-wirtschafts-nachrichten.de",
		IpAddress: "198.41.185.52",
	},
	&fronted.Masquerade{
		Domain:    "diablofans.com",
		IpAddress: "198.41.209.103",
	},
	&fronted.Masquerade{
		Domain:    "dnevnik.hr",
		IpAddress: "190.93.241.21",
	},
	&fronted.Masquerade{
		Domain:    "dostor.org",
		IpAddress: "104.20.12.195",
	},
	&fronted.Masquerade{
		Domain:    "downloadming.nu",
		IpAddress: "198.41.189.76",
	},
	&fronted.Masquerade{
		Domain:    "dpstream.net",
		IpAddress: "198.41.191.151",
	},
	&fronted.Masquerade{
		Domain:    "drakulastream.eu",
		IpAddress: "141.101.125.249",
	},
	&fronted.Masquerade{
		Domain:    "drp.su",
		IpAddress: "162.159.244.17",
	},
	&fronted.Masquerade{
		Domain:    "dumpaday.com",
		IpAddress: "162.159.242.119",
	},
	&fronted.Masquerade{
		Domain:    "e-cigarette-forum.com",
		IpAddress: "104.20.30.178",
	},
	&fronted.Masquerade{
		Domain:    "e-monsite.com",
		IpAddress: "141.101.120.122",
	},
	&fronted.Masquerade{
		Domain:    "e-radio.gr",
		IpAddress: "198.41.183.19",
	},
	&fronted.Masquerade{
		Domain:    "eclypsia.com",
		IpAddress: "141.101.112.98",
	},
	&fronted.Masquerade{
		Domain:    "edublogs.org",
		IpAddress: "104.16.1.23",
	},
	&fronted.Masquerade{
		Domain:    "egaliteetreconciliation.fr",
		IpAddress: "190.93.243.80",
	},
	&fronted.Masquerade{
		Domain:    "egyup.com",
		IpAddress: "108.162.201.33",
	},
	&fronted.Masquerade{
		Domain:    "eharmony.com",
		IpAddress: "199.83.133.3",
	},
	&fronted.Masquerade{
		Domain:    "einthusan.com",
		IpAddress: "198.41.190.126",
	},
	&fronted.Masquerade{
		Domain:    "elhacker.net",
		IpAddress: "108.162.205.73",
	},
	&fronted.Masquerade{
		Domain:    "elwatannews.com",
		IpAddress: "104.20.15.106",
	},
	&fronted.Masquerade{
		Domain:    "en.bitcoin.it",
		IpAddress: "162.159.246.241",
	},
	&fronted.Masquerade{
		Domain:    "eslamoda.com",
		IpAddress: "162.159.253.119",
	},
	&fronted.Masquerade{
		Domain:    "etorrent.co.kr",
		IpAddress: "198.41.190.119",
	},
	&fronted.Masquerade{
		Domain:    "etvnet.com",
		IpAddress: "104.20.2.29",
	},
	&fronted.Masquerade{
		Domain:    "euw.leagueoflegends.com",
		IpAddress: "104.16.21.33",
	},
	&fronted.Masquerade{
		Domain:    "evozi.com",
		IpAddress: "198.41.203.14",
	},
	&fronted.Masquerade{
		Domain:    "explosm.net",
		IpAddress: "198.41.205.239",
	},
	&fronted.Masquerade{
		Domain:    "extratorrent.cc",
		IpAddress: "162.159.254.81",
	},
	&fronted.Masquerade{
		Domain:    "faithtap.com",
		IpAddress: "198.41.189.57",
	},
	&fronted.Masquerade{
		Domain:    "famousbirthdays.com",
		IpAddress: "190.93.245.80",
	},
	&fronted.Masquerade{
		Domain:    "fasttech.com",
		IpAddress: "190.93.241.97",
	},
	&fronted.Masquerade{
		Domain:    "feedly.com",
		IpAddress: "162.159.254.4",
	},
	&fronted.Masquerade{
		Domain:    "filesfetcher.com",
		IpAddress: "198.41.185.168",
	},
	&fronted.Masquerade{
		Domain:    "filmesonlinegratis.net",
		IpAddress: "141.101.113.38",
	},
	&fronted.Masquerade{
		Domain:    "fiverr.com",
		IpAddress: "192.33.31.61",
	},
	&fronted.Masquerade{
		Domain:    "flashgames.it",
		IpAddress: "141.101.120.119",
	},
	&fronted.Masquerade{
		Domain:    "food52.com",
		IpAddress: "104.20.0.127",
	},
	&fronted.Masquerade{
		Domain:    "forexpeacearmy.com",
		IpAddress: "190.93.241.28",
	},
	&fronted.Masquerade{
		Domain:    "forgifs.com",
		IpAddress: "162.159.251.66",
	},
	&fronted.Masquerade{
		Domain:    "freedoge.co.in",
		IpAddress: "141.101.127.23",
	},
	&fronted.Masquerade{
		Domain:    "freenode.net",
		IpAddress: "162.159.249.27",
	},
	&fronted.Masquerade{
		Domain:    "frontpage.fok.nl",
		IpAddress: "104.20.11.180",
	},
	&fronted.Masquerade{
		Domain:    "fshare.vn",
		IpAddress: "118.69.164.133",
	},
	&fronted.Masquerade{
		Domain:    "fsplay.net",
		IpAddress: "198.41.247.238",
	},
	&fronted.Masquerade{
		Domain:    "fun698.com",
		IpAddress: "198.41.206.118",
	},
	&fronted.Masquerade{
		Domain:    "funnymama.com",
		IpAddress: "198.41.249.64",
	},
	&fronted.Masquerade{
		Domain:    "futhead.com",
		IpAddress: "141.101.115.99",
	},
	&fronted.Masquerade{
		Domain:    "gahe.com",
		IpAddress: "162.159.252.233",
	},
	&fronted.Masquerade{
		Domain:    "gamebaby.com",
		IpAddress: "162.159.241.107",
	},
	&fronted.Masquerade{
		Domain:    "gameninja.com",
		IpAddress: "198.41.188.17",
	},
	&fronted.Masquerade{
		Domain:    "gamepedia.com",
		IpAddress: "141.101.114.101",
	},
	&fronted.Masquerade{
		Domain:    "gameskwala.com",
		IpAddress: "162.159.241.227",
	},
	&fronted.Masquerade{
		Domain:    "gcflearnfree.org",
		IpAddress: "141.101.112.72",
	},
	&fronted.Masquerade{
		Domain:    "geo.tv",
		IpAddress: "190.93.244.11",
	},
	&fronted.Masquerade{
		Domain:    "gfycat.com",
		IpAddress: "198.41.209.26",
	},
	&fronted.Masquerade{
		Domain:    "ghost.org",
		IpAddress: "190.93.245.19",
	},
	&fronted.Masquerade{
		Domain:    "gigacircle.com",
		IpAddress: "104.16.0.36",
	},
	&fronted.Masquerade{
		Domain:    "gilt.com",
		IpAddress: "198.41.208.113",
	},
	&fronted.Masquerade{
		Domain:    "gizmodo.com.au",
		IpAddress: "141.101.114.74",
	},
	&fronted.Masquerade{
		Domain:    "glassdoor.com",
		IpAddress: "190.93.244.224",
	},
	&fronted.Masquerade{
		Domain:    "globalresearch.ca",
		IpAddress: "162.159.246.162",
	},
	&fronted.Masquerade{
		Domain:    "gooddrama.net",
		IpAddress: "198.41.204.151",
	},
	&fronted.Masquerade{
		Domain:    "goodsearch.com",
		IpAddress: "190.93.241.98",
	},
	&fronted.Masquerade{
		Domain:    "gooool.org",
		IpAddress: "162.159.243.194",
	},
	&fronted.Masquerade{
		Domain:    "gosugamers.net",
		IpAddress: "162.159.240.238",
	},
	&fronted.Masquerade{
		Domain:    "gottabemobile.com",
		IpAddress: "190.93.241.110",
	},
	&fronted.Masquerade{
		Domain:    "goud.ma",
		IpAddress: "141.101.126.204",
	},
	&fronted.Masquerade{
		Domain:    "gtspirit.com",
		IpAddress: "162.159.243.151",
	},
	&fronted.Masquerade{
		Domain:    "guardianlv.com",
		IpAddress: "198.41.204.133",
	},
	&fronted.Masquerade{
		Domain:    "haber1903.com",
		IpAddress: "141.101.126.111",
	},
	&fronted.Masquerade{
		Domain:    "hackforums.net",
		IpAddress: "141.101.121.10",
	},
	&fronted.Masquerade{
		Domain:    "hardmob.com.br",
		IpAddress: "190.93.241.96",
	},
	&fronted.Masquerade{
		Domain:    "hearthpwn.com",
		IpAddress: "190.93.244.113",
	},
	&fronted.Masquerade{
		Domain:    "hitleap.com",
		IpAddress: "198.41.183.88",
	},
	&fronted.Masquerade{
		Domain:    "hobbyking.com",
		IpAddress: "190.93.243.124",
	},
	&fronted.Masquerade{
		Domain:    "home.ijreview.com",
		IpAddress: "104.16.2.43",
	},
	&fronted.Masquerade{
		Domain:    "i-fit.com.tw",
		IpAddress: "108.162.202.108",
	},
	&fronted.Masquerade{
		Domain:    "ibuildapp.com",
		IpAddress: "141.101.112.201",
	},
	&fronted.Masquerade{
		Domain:    "ifilez.org",
		IpAddress: "190.93.240.94",
	},
	&fronted.Masquerade{
		Domain:    "iitv.info",
		IpAddress: "162.159.251.221",
	},
	&fronted.Masquerade{
		Domain:    "ikman.lk",
		IpAddress: "104.16.19.214",
	},
	&fronted.Masquerade{
		Domain:    "imgchili.net",
		IpAddress: "198.41.207.163",
	},
	&fronted.Masquerade{
		Domain:    "imgflip.com",
		IpAddress: "190.93.245.142",
	},
	&fronted.Masquerade{
		Domain:    "imgspice.com",
		IpAddress: "198.41.249.212",
	},
	&fronted.Masquerade{
		Domain:    "index.hr",
		IpAddress: "198.41.182.4",
	},
	&fronted.Masquerade{
		Domain:    "inforesist.org",
		IpAddress: "108.162.206.29",
	},
	&fronted.Masquerade{
		Domain:    "intercambiosvirtuales.org",
		IpAddress: "162.159.242.146",
	},
	&fronted.Masquerade{
		Domain:    "ipiccy.com",
		IpAddress: "190.93.241.33",
	},
	&fronted.Masquerade{
		Domain:    "iplocation.net",
		IpAddress: "104.20.16.186",
	},
	&fronted.Masquerade{
		Domain:    "iptorrents.com",
		IpAddress: "141.101.123.132",
	},
	&fronted.Masquerade{
		Domain:    "isohunt.to",
		IpAddress: "198.41.190.233",
	},
	&fronted.Masquerade{
		Domain:    "israelvideonetwork.com",
		IpAddress: "198.41.191.72",
	},
	&fronted.Masquerade{
		Domain:    "italia-film.org",
		IpAddress: "190.93.242.91",
	},
	&fronted.Masquerade{
		Domain:    "iwebchk.com",
		IpAddress: "162.159.241.191",
	},
	&fronted.Masquerade{
		Domain:    "ixl.com",
		IpAddress: "141.101.115.137",
	},
	&fronted.Masquerade{
		Domain:    "j.gs",
		IpAddress: "198.41.249.35",
	},
	&fronted.Masquerade{
		Domain:    "jamiiforums.com",
		IpAddress: "162.159.242.71",
	},
	&fronted.Masquerade{
		Domain:    "jquery.com",
		IpAddress: "104.16.14.15",
	},
	&fronted.Masquerade{
		Domain:    "jquerymobile.com",
		IpAddress: "104.16.10.13",
	},
	&fronted.Masquerade{
		Domain:    "jqueryui.com",
		IpAddress: "104.16.2.14",
	},
	&fronted.Masquerade{
		Domain:    "jumia.com.ng",
		IpAddress: "198.41.188.216",
	},
	&fronted.Masquerade{
		Domain:    "k2s.cc",
		IpAddress: "162.159.244.42",
	},
	&fronted.Masquerade{
		Domain:    "karatbars.com",
		IpAddress: "205.204.91.35",
	},
	&fronted.Masquerade{
		Domain:    "karnaval.com",
		IpAddress: "141.101.121.196",
	},
	&fronted.Masquerade{
		Domain:    "kaymu.com.ng",
		IpAddress: "104.20.26.2",
	},
	&fronted.Masquerade{
		Domain:    "kaymu.pk",
		IpAddress: "141.101.113.175",
	},
	&fronted.Masquerade{
		Domain:    "kinogo.net",
		IpAddress: "141.101.123.114",
	},
	&fronted.Masquerade{
		Domain:    "kinoman.tv",
		IpAddress: "198.41.188.174",
	},
	&fronted.Masquerade{
		Domain:    "klix.ba",
		IpAddress: "190.93.242.87",
	},
	&fronted.Masquerade{
		Domain:    "korben.info",
		IpAddress: "162.159.251.186",
	},
	&fronted.Masquerade{
		Domain:    "kwejk.pl",
		IpAddress: "162.159.254.38",
	},
	&fronted.Masquerade{
		Domain:    "lamido.co.id",
		IpAddress: "198.41.187.224",
	},
	&fronted.Masquerade{
		Domain:    "lapatilla.com",
		IpAddress: "141.101.113.240",
	},
	&fronted.Masquerade{
		Domain:    "lasvegassun.com",
		IpAddress: "141.101.123.129",
	},
	&fronted.Masquerade{
		Domain:    "laughingsquid.com",
		IpAddress: "162.159.247.97",
	},
	&fronted.Masquerade{
		Domain:    "lbcgroup.tv",
		IpAddress: "141.101.112.51",
	},
	&fronted.Masquerade{
		Domain:    "legacyclix.com",
		IpAddress: "162.159.249.65",
	},
	&fronted.Masquerade{
		Domain:    "legiaodosherois.com.br",
		IpAddress: "198.41.205.241",
	},
	&fronted.Masquerade{
		Domain:    "libertyland.tv",
		IpAddress: "104.20.1.179",
	},
	&fronted.Masquerade{
		Domain:    "lifebuzz.com",
		IpAddress: "104.16.21.166",
	},
	&fronted.Masquerade{
		Domain:    "lifehacker.com.au",
		IpAddress: "190.93.247.73",
	},
	&fronted.Masquerade{
		Domain:    "likemag.com",
		IpAddress: "162.159.251.215",
	},
	&fronted.Masquerade{
		Domain:    "likes.com",
		IpAddress: "190.93.244.34",
	},
	&fronted.Masquerade{
		Domain:    "listenpersian.net",
		IpAddress: "198.41.249.9",
	},
	&fronted.Masquerade{
		Domain:    "localbitcoins.com",
		IpAddress: "104.20.31.3",
	},
	&fronted.Masquerade{
		Domain:    "maannews.net",
		IpAddress: "198.41.180.195",
	},
	&fronted.Masquerade{
		Domain:    "macacovelho.com.br",
		IpAddress: "198.41.188.108",
	},
	&fronted.Masquerade{
		Domain:    "macworld.co.uk",
		IpAddress: "104.16.12.54",
	},
	&fronted.Masquerade{
		Domain:    "madmimi.com",
		IpAddress: "141.101.123.192",
	},
	&fronted.Masquerade{
		Domain:    "mafiashare.net",
		IpAddress: "141.101.120.96",
	},
	&fronted.Masquerade{
		Domain:    "manicomio-share.com",
		IpAddress: "198.41.207.93",
	},
	&fronted.Masquerade{
		Domain:    "maplestage.com",
		IpAddress: "162.159.254.194",
	},
	&fronted.Masquerade{
		Domain:    "marunadanmalayali.com",
		IpAddress: "141.101.126.226",
	},
	&fronted.Masquerade{
		Domain:    "matchesfashion.com",
		IpAddress: "198.41.185.14",
	},
	&fronted.Masquerade{
		Domain:    "mazika2day.com",
		IpAddress: "141.101.127.188",
	},
	&fronted.Masquerade{
		Domain:    "media-fire.org",
		IpAddress: "198.41.188.89",
	},
	&fronted.Masquerade{
		Domain:    "microworkers.com",
		IpAddress: "190.93.243.147",
	},
	&fronted.Masquerade{
		Domain:    "minecraftforum.net",
		IpAddress: "190.93.247.117",
	},
	&fronted.Masquerade{
		Domain:    "minecraftservers.org",
		IpAddress: "190.93.243.14",
	},
	&fronted.Masquerade{
		Domain:    "mixedmartialarts.com",
		IpAddress: "141.101.113.57",
	},
	&fronted.Masquerade{
		Domain:    "mixergy.com",
		IpAddress: "162.159.251.147",
	},
	&fronted.Masquerade{
		Domain:    "mmo-champion.com",
		IpAddress: "190.93.244.118",
	},
	&fronted.Masquerade{
		Domain:    "mo.gov",
		IpAddress: "104.16.23.39",
	},
	&fronted.Masquerade{
		Domain:    "mobafire.com",
		IpAddress: "141.101.121.19",
	},
	&fronted.Masquerade{
		Domain:    "modern.az",
		IpAddress: "108.162.206.159",
	},
	&fronted.Masquerade{
		Domain:    "moneyplatform.biz",
		IpAddress: "108.162.201.215",
	},
	&fronted.Masquerade{
		Domain:    "my-hit.org",
		IpAddress: "104.20.4.121",
	},
	&fronted.Masquerade{
		Domain:    "mylikes.com",
		IpAddress: "190.93.247.37",
	},
	&fronted.Masquerade{
		Domain:    "naijaloaded.com.ng",
		IpAddress: "141.101.126.140",
	},
	&fronted.Masquerade{
		Domain:    "nairaland.com",
		IpAddress: "198.41.185.68",
	},
	&fronted.Masquerade{
		Domain:    "naosalvo.com.br",
		IpAddress: "141.101.124.184",
	},
	&fronted.Masquerade{
		Domain:    "nationalreview.com",
		IpAddress: "190.93.247.199",
	},
	&fronted.Masquerade{
		Domain:    "nbadraft.net",
		IpAddress: "162.159.251.170",
	},
	&fronted.Masquerade{
		Domain:    "network-tools.com",
		IpAddress: "141.101.123.110",
	},
	&fronted.Masquerade{
		Domain:    "newmobilelife.com",
		IpAddress: "108.162.207.54",
	},
	&fronted.Masquerade{
		Domain:    "nextmedia.com",
		IpAddress: "104.16.9.5",
	},
	&fronted.Masquerade{
		Domain:    "noticiaaldia.com",
		IpAddress: "198.41.180.92",
	},
	&fronted.Masquerade{
		Domain:    "noticierodigital.com",
		IpAddress: "190.93.241.95",
	},
	&fronted.Masquerade{
		Domain:    "ojooo.com",
		IpAddress: "198.41.190.230",
	},
	&fronted.Masquerade{
		Domain:    "omgtorrent.com",
		IpAddress: "198.41.202.97",
	},
	&fronted.Masquerade{
		Domain:    "oneplus.net",
		IpAddress: "141.101.126.10",
	},
	&fronted.Masquerade{
		Domain:    "onlineclock.net",
		IpAddress: "190.93.241.58",
	},
	&fronted.Masquerade{
		Domain:    "onlinesoccermanager.com",
		IpAddress: "162.159.253.18",
	},
	&fronted.Masquerade{
		Domain:    "opencart.com",
		IpAddress: "104.20.15.19",
	},
	&fronted.Masquerade{
		Domain:    "opensoftwareupdater.com",
		IpAddress: "198.41.209.253",
	},
	&fronted.Masquerade{
		Domain:    "opposingviews.com",
		IpAddress: "162.159.253.200",
	},
	&fronted.Masquerade{
		Domain:    "oscaro.com",
		IpAddress: "104.16.12.97",
	},
	&fronted.Masquerade{
		Domain:    "osdir.com",
		IpAddress: "162.159.252.186",
	},
	&fronted.Masquerade{
		Domain:    "palemoon.org",
		IpAddress: "162.159.254.211",
	},
	&fronted.Masquerade{
		Domain:    "pastebin.com",
		IpAddress: "190.93.240.15",
	},
	&fronted.Masquerade{
		Domain:    "pcadvisor.co.uk",
		IpAddress: "104.16.26.51",
	},
	&fronted.Masquerade{
		Domain:    "pelis24.com",
		IpAddress: "198.41.188.143",
	},
	&fronted.Masquerade{
		Domain:    "photoyoum7.com",
		IpAddress: "104.16.7.117",
	},
	&fronted.Masquerade{
		Domain:    "pijamasurf.com",
		IpAddress: "162.159.241.249",
	},
	&fronted.Masquerade{
		Domain:    "planetminecraft.com",
		IpAddress: "141.101.123.126",
	},
	&fronted.Masquerade{
		Domain:    "plp.cl",
		IpAddress: "198.41.200.28",
	},
	&fronted.Masquerade{
		Domain:    "podomatic.com",
		IpAddress: "104.20.20.4",
	},
	&fronted.Masquerade{
		Domain:    "popcash.net",
		IpAddress: "198.41.206.216",
	},
	&fronted.Masquerade{
		Domain:    "popnhop.com",
		IpAddress: "162.159.248.205",
	},
	&fronted.Masquerade{
		Domain:    "post852.com",
		IpAddress: "104.20.17.192",
	},
	&fronted.Masquerade{
		Domain:    "postcron.com",
		IpAddress: "162.159.242.38",
	},
	&fronted.Masquerade{
		Domain:    "postto.me",
		IpAddress: "141.101.120.157",
	},
	&fronted.Masquerade{
		Domain:    "premium.wpmudev.org",
		IpAddress: "104.16.24.10",
	},
	&fronted.Masquerade{
		Domain:    "premiumbeat.com",
		IpAddress: "104.20.2.94",
	},
	&fronted.Masquerade{
		Domain:    "premiumwp.com",
		IpAddress: "162.159.250.100",
	},
	&fronted.Masquerade{
		Domain:    "prlog.ru",
		IpAddress: "162.159.242.63",
	},
	&fronted.Masquerade{
		Domain:    "prntscr.com",
		IpAddress: "198.41.191.131",
	},
	&fronted.Masquerade{
		Domain:    "proboards.com",
		IpAddress: "190.93.244.205",
	},
	&fronted.Masquerade{
		Domain:    "proprofs.com",
		IpAddress: "198.41.207.245",
	},
	&fronted.Masquerade{
		Domain:    "proteusthemes.com",
		IpAddress: "162.159.248.215",
	},
	&fronted.Masquerade{
		Domain:    "proxybay.info",
		IpAddress: "91.236.116.247",
	},
	&fronted.Masquerade{
		Domain:    "ptcstair.com",
		IpAddress: "162.159.250.107",
	},
	&fronted.Masquerade{
		Domain:    "qafqazinfo.az",
		IpAddress: "162.159.244.58",
	},
	&fronted.Masquerade{
		Domain:    "qatarliving.com",
		IpAddress: "198.41.249.175",
	},
	&fronted.Masquerade{
		Domain:    "qol.az",
		IpAddress: "162.159.243.133",
	},
	&fronted.Masquerade{
		Domain:    "r10.net",
		IpAddress: "104.20.27.135",
	},
	&fronted.Masquerade{
		Domain:    "rapgenius.com",
		IpAddress: "104.16.26.4",
	},
	&fronted.Masquerade{
		Domain:    "rapradar.com",
		IpAddress: "141.101.112.16",
	},
	&fronted.Masquerade{
		Domain:    "rassd.com",
		IpAddress: "162.159.254.222",
	},
	&fronted.Masquerade{
		Domain:    "reshareworthy.com",
		IpAddress: "141.101.127.122",
	},
	&fronted.Masquerade{
		Domain:    "rollingout.com",
		IpAddress: "198.41.187.117",
	},
	&fronted.Masquerade{
		Domain:    "rsw-systems.com",
		IpAddress: "104.20.18.116",
	},
	&fronted.Masquerade{
		Domain:    "rudaw.net",
		IpAddress: "190.93.240.83",
	},
	&fronted.Masquerade{
		Domain:    "rus.ec",
		IpAddress: "198.41.186.201",
	},
	&fronted.Masquerade{
		Domain:    "rusvesna.su",
		IpAddress: "162.159.253.5",
	},
	&fronted.Masquerade{
		Domain:    "sa.ae",
		IpAddress: "162.159.240.111",
	},
	&fronted.Masquerade{
		Domain:    "sabq.org",
		IpAddress: "190.93.244.116",
	},
	&fronted.Masquerade{
		Domain:    "sanakirja.org",
		IpAddress: "190.93.241.90",
	},
	&fronted.Masquerade{
		Domain:    "sayidaty.net",
		IpAddress: "108.162.202.30",
	},
	&fronted.Masquerade{
		Domain:    "scotch.io",
		IpAddress: "141.101.125.86",
	},
	&fronted.Masquerade{
		Domain:    "searchengines.guru",
		IpAddress: "190.93.241.113",
	},
	&fronted.Masquerade{
		Domain:    "searchengines.ru",
		IpAddress: "190.93.241.113",
	},
	&fronted.Masquerade{
		Domain:    "seemorgh.com",
		IpAddress: "141.101.120.194",
	},
	&fronted.Masquerade{
		Domain:    "sendgrid.com",
		IpAddress: "104.20.20.26",
	},
	&fronted.Masquerade{
		Domain:    "sergey-mavrodi-mmm.org",
		IpAddress: "141.101.125.231",
	},
	&fronted.Masquerade{
		Domain:    "sergey-mavrodi.com",
		IpAddress: "104.20.4.194",
	},
	&fronted.Masquerade{
		Domain:    "sergeymavrodi.com",
		IpAddress: "104.20.30.193",
	},
	&fronted.Masquerade{
		Domain:    "shahiya.com",
		IpAddress: "162.159.240.128",
	},
	&fronted.Masquerade{
		Domain:    "shapeways.com",
		IpAddress: "198.41.189.36",
	},
	&fronted.Masquerade{
		Domain:    "sheknows.com",
		IpAddress: "190.93.249.19",
	},
	&fronted.Masquerade{
		Domain:    "shmoop.com",
		IpAddress: "204.232.154.6",
	},
	&fronted.Masquerade{
		Domain:    "siam-movie.com",
		IpAddress: "198.41.181.78",
	},
	&fronted.Masquerade{
		Domain:    "siliconera.com",
		IpAddress: "190.93.246.99",
	},
	&fronted.Masquerade{
		Domain:    "siliconrus.com",
		IpAddress: "198.41.191.66",
	},
	&fronted.Masquerade{
		Domain:    "sinchew.com.my",
		IpAddress: "141.101.121.131",
	},
	&fronted.Masquerade{
		Domain:    "sitetalk.com",
		IpAddress: "190.93.241.207",
	},
	&fronted.Masquerade{
		Domain:    "skladchik.com",
		IpAddress: "104.20.3.89",
	},
	&fronted.Masquerade{
		Domain:    "smallpdf.com",
		IpAddress: "107.170.126.92",
	},
	&fronted.Masquerade{
		Domain:    "smittenkitchen.com",
		IpAddress: "141.101.123.139",
	},
	&fronted.Masquerade{
		Domain:    "smosh.com",
		IpAddress: "162.159.254.34",
	},
	&fronted.Masquerade{
		Domain:    "smotrisport.tv",
		IpAddress: "198.41.182.22",
	},
	&fronted.Masquerade{
		Domain:    "snapengage.com",
		IpAddress: "141.101.112.133",
	},
	&fronted.Masquerade{
		Domain:    "snip.ly",
		IpAddress: "108.162.201.204",
	},
	&fronted.Masquerade{
		Domain:    "softarchive.net",
		IpAddress: "108.162.202.222",
	},
	&fronted.Masquerade{
		Domain:    "somuch.com",
		IpAddress: "141.101.127.228",
	},
	&fronted.Masquerade{
		Domain:    "songspk.name",
		IpAddress: "108.162.201.183",
	},
	&fronted.Masquerade{
		Domain:    "soompi.com",
		IpAddress: "104.20.19.19",
	},
	&fronted.Masquerade{
		Domain:    "sooperarticles.com",
		IpAddress: "108.162.206.236",
	},
	&fronted.Masquerade{
		Domain:    "spi0n.com",
		IpAddress: "198.41.180.58",
	},
	&fronted.Masquerade{
		Domain:    "sportbox.az",
		IpAddress: "108.162.206.207",
	},
	&fronted.Masquerade{
		Domain:    "sprotyv.info",
		IpAddress: "141.101.126.17",
	},
	&fronted.Masquerade{
		Domain:    "stadt-bremerhaven.de",
		IpAddress: "198.41.186.16",
	},
	&fronted.Masquerade{
		Domain:    "stagram.com",
		IpAddress: "190.93.242.45",
	},
	&fronted.Masquerade{
		Domain:    "stansberryresearch.com",
		IpAddress: "104.20.27.17",
	},
	&fronted.Masquerade{
		Domain:    "statcounter.com",
		IpAddress: "104.20.3.47",
	},
	&fronted.Masquerade{
		Domain:    "steamdb.info",
		IpAddress: "162.159.254.176",
	},
	&fronted.Masquerade{
		Domain:    "streamallthis.me",
		IpAddress: "162.159.241.171",
	},
	&fronted.Masquerade{
		Domain:    "subscene.com",
		IpAddress: "190.93.255.42",
	},
	&fronted.Masquerade{
		Domain:    "sudaneseonline.com",
		IpAddress: "198.41.206.254",
	},
	&fronted.Masquerade{
		Domain:    "super.ae",
		IpAddress: "162.159.255.6",
	},
	&fronted.Masquerade{
		Domain:    "suzukikenichi.com",
		IpAddress: "157.112.184.64",
	},
	&fronted.Masquerade{
		Domain:    "t24.com.tr",
		IpAddress: "141.101.125.79",
	},
	&fronted.Masquerade{
		Domain:    "tahrirnews.com",
		IpAddress: "198.41.191.205",
	},
	&fronted.Masquerade{
		Domain:    "tarafdari.com",
		IpAddress: "198.41.191.174",
	},
	&fronted.Masquerade{
		Domain:    "tech-recipes.com",
		IpAddress: "141.101.124.86",
	},
	&fronted.Masquerade{
		Domain:    "teebik.com",
		IpAddress: "198.41.189.108",
	},
	&fronted.Masquerade{
		Domain:    "templatemonster.com",
		IpAddress: "104.20.29.119",
	},
	&fronted.Masquerade{
		Domain:    "teveonline.net",
		IpAddress: "198.41.205.151",
	},
	&fronted.Masquerade{
		Domain:    "tgju.org",
		IpAddress: "108.162.207.230",
	},
	&fronted.Masquerade{
		Domain:    "the-open-mind.com",
		IpAddress: "141.101.126.65",
	},
	&fronted.Masquerade{
		Domain:    "thediplomat.com",
		IpAddress: "162.159.240.235",
	},
	&fronted.Masquerade{
		Domain:    "thefile.me",
		IpAddress: "162.159.252.33",
	},
	&fronted.Masquerade{
		Domain:    "thefreethoughtproject.com",
		IpAddress: "162.159.240.157",
	},
	&fronted.Masquerade{
		Domain:    "theiconic.com.au",
		IpAddress: "198.41.186.186",
	},
	&fronted.Masquerade{
		Domain:    "theladbible.com",
		IpAddress: "198.41.215.5",
	},
	&fronted.Masquerade{
		Domain:    "themattwalshblog.com",
		IpAddress: "108.162.203.50",
	},
	&fronted.Masquerade{
		Domain:    "theme-fusion.com",
		IpAddress: "104.20.27.96",
	},
	&fronted.Masquerade{
		Domain:    "themindunleashed.org",
		IpAddress: "108.162.201.137",
	},
	&fronted.Masquerade{
		Domain:    "thenationonlineng.net",
		IpAddress: "162.159.254.179",
	},
	&fronted.Masquerade{
		Domain:    "thenews.com.pk",
		IpAddress: "104.16.30.122",
	},
	&fronted.Masquerade{
		Domain:    "thenewstribe.com",
		IpAddress: "162.159.246.82",
	},
	&fronted.Masquerade{
		Domain:    "thepioneerwoman.com",
		IpAddress: "198.41.190.137",
	},
	&fronted.Masquerade{
		Domain:    "therakyatpost.com",
		IpAddress: "198.41.187.177",
	},
	&fronted.Masquerade{
		Domain:    "thesportbible.com",
		IpAddress: "141.101.114.97",
	},
	&fronted.Masquerade{
		Domain:    "thevideo.me",
		IpAddress: "162.159.242.240",
	},
	&fronted.Masquerade{
		Domain:    "thisiswhyimbroke.com",
		IpAddress: "162.159.250.214",
	},
	&fronted.Masquerade{
		Domain:    "tickld.com",
		IpAddress: "104.16.27.6",
	},
	&fronted.Masquerade{
		Domain:    "todayifoundout.com",
		IpAddress: "141.101.126.98",
	},
	&fronted.Masquerade{
		Domain:    "torlock.com",
		IpAddress: "198.41.200.25",
	},
	&fronted.Masquerade{
		Domain:    "torrentleech.org",
		IpAddress: "108.162.200.95",
	},
	&fronted.Masquerade{
		Domain:    "trafficgenesis.com",
		IpAddress: "162.159.242.133",
	},
	&fronted.Masquerade{
		Domain:    "tribalfootball.com",
		IpAddress: "141.101.113.4",
	},
	&fronted.Masquerade{
		Domain:    "tripleclicks.com",
		IpAddress: "141.101.125.214",
	},
	&fronted.Masquerade{
		Domain:    "tructiepbongda.com",
		IpAddress: "198.41.201.45",
	},
	&fronted.Masquerade{
		Domain:    "trueactivist.com",
		IpAddress: "162.159.255.134",
	},
	&fronted.Masquerade{
		Domain:    "tutsplus.com",
		IpAddress: "190.93.242.15",
	},
	&fronted.Masquerade{
		Domain:    "twentytwowords.com",
		IpAddress: "162.159.244.22",
	},
	&fronted.Masquerade{
		Domain:    "udemy.com",
		IpAddress: "141.101.113.23",
	},
	&fronted.Masquerade{
		Domain:    "ummat.net",
		IpAddress: "108.162.207.43",
	},
	&fronted.Masquerade{
		Domain:    "uniladmag.com",
		IpAddress: "198.41.206.219",
	},
	&fronted.Masquerade{
		Domain:    "unwire.hk",
		IpAddress: "198.41.188.172",
	},
	&fronted.Masquerade{
		Domain:    "updatenowpro.com",
		IpAddress: "162.159.243.60",
	},
	&fronted.Masquerade{
		Domain:    "updatersoft.com",
		IpAddress: "162.159.248.72",
	},
	&fronted.Masquerade{
		Domain:    "uploadboy.com",
		IpAddress: "141.101.125.9",
	},
	&fronted.Masquerade{
		Domain:    "uptimerobot.com",
		IpAddress: "174.36.49.98",
	},
	&fronted.Masquerade{
		Domain:    "uptobox.com",
		IpAddress: "141.101.112.96",
	},
	&fronted.Masquerade{
		Domain:    "urbanfonts.com",
		IpAddress: "162.159.240.64",
	},
	&fronted.Masquerade{
		Domain:    "verseriesynovelas.com",
		IpAddress: "141.101.125.48",
	},
	&fronted.Masquerade{
		Domain:    "vertele.com",
		IpAddress: "162.159.244.94",
	},
	&fronted.Masquerade{
		Domain:    "videomega.tv",
		IpAddress: "162.159.253.155",
	},
	&fronted.Masquerade{
		Domain:    "videostripe.com",
		IpAddress: "198.41.187.157",
	},
	&fronted.Masquerade{
		Domain:    "videoyoum7.com",
		IpAddress: "104.16.25.116",
	},
	&fronted.Masquerade{
		Domain:    "vitorrent.org",
		IpAddress: "162.159.244.211",
	},
	&fronted.Masquerade{
		Domain:    "vladtv.com",
		IpAddress: "162.159.254.30",
	},
	&fronted.Masquerade{
		Domain:    "vodlocker.com",
		IpAddress: "162.159.246.224",
	},
	&fronted.Masquerade{
		Domain:    "vodly.to",
		IpAddress: "190.93.242.35",
	},
	&fronted.Masquerade{
		Domain:    "voetbalzone.nl",
		IpAddress: "198.41.191.199",
	},
	&fronted.Masquerade{
		Domain:    "vr-zone.com",
		IpAddress: "162.159.251.175",
	},
	&fronted.Masquerade{
		Domain:    "watchfreemovies.ch",
		IpAddress: "190.93.240.21",
	},
	&fronted.Masquerade{
		Domain:    "watchseries-online.ch",
		IpAddress: "104.20.29.123",
	},
	&fronted.Masquerade{
		Domain:    "watchserieshd.eu",
		IpAddress: "104.20.9.121",
	},
	&fronted.Masquerade{
		Domain:    "weknowmemes.com",
		IpAddress: "162.159.252.35",
	},
	&fronted.Masquerade{
		Domain:    "what.cd",
		IpAddress: "198.41.190.106",
	},
	&fronted.Masquerade{
		Domain:    "whatculture.com",
		IpAddress: "198.41.249.80",
	},
	&fronted.Masquerade{
		Domain:    "wholehk.com",
		IpAddress: "198.41.205.227",
	},
	&fronted.Masquerade{
		Domain:    "wikiwiki.jp",
		IpAddress: "190.93.240.68",
	},
	&fronted.Masquerade{
		Domain:    "wiziq.com",
		IpAddress: "190.93.244.247",
	},
	&fronted.Masquerade{
		Domain:    "woorank.com",
		IpAddress: "54.165.180.233",
	},
	&fronted.Masquerade{
		Domain:    "www.4chan.org",
		IpAddress: "141.101.114.6",
	},
	&fronted.Masquerade{
		Domain:    "www.aciprensa.com",
		IpAddress: "198.41.188.166",
	},
	&fronted.Masquerade{
		Domain:    "www.addtoany.com",
		IpAddress: "141.101.125.160",
	},
	&fronted.Masquerade{
		Domain:    "www.alweeam.com.sa",
		IpAddress: "141.101.126.49",
	},
	&fronted.Masquerade{
		Domain:    "www.animenewsnetwork.com",
		IpAddress: "198.41.183.80",
	},
	&fronted.Masquerade{
		Domain:    "www.autostraddle.com",
		IpAddress: "162.159.247.115",
	},
	&fronted.Masquerade{
		Domain:    "www.binary.com",
		IpAddress: "190.93.241.81",
	},
	&fronted.Masquerade{
		Domain:    "www.bj2.me",
		IpAddress: "190.93.240.108",
	},
	&fronted.Masquerade{
		Domain:    "www.brasil247.com",
		IpAddress: "162.159.251.62",
	},
	&fronted.Masquerade{
		Domain:    "www.bulletproofexec.com",
		IpAddress: "104.20.11.19",
	},
	&fronted.Masquerade{
		Domain:    "www.burnews.com",
		IpAddress: "190.93.243.102",
	},
	&fronted.Masquerade{
		Domain:    "www.cairodar.com",
		IpAddress: "104.16.29.116",
	},
	&fronted.Masquerade{
		Domain:    "www.campingworld.com",
		IpAddress: "141.101.112.210",
	},
	&fronted.Masquerade{
		Domain:    "www.caracoltv.com",
		IpAddress: "190.93.240.64",
	},
	&fronted.Masquerade{
		Domain:    "www.cbox.ws",
		IpAddress: "198.11.203.46",
	},
	&fronted.Masquerade{
		Domain:    "www.change.org",
		IpAddress: "104.16.5.13",
	},
	&fronted.Masquerade{
		Domain:    "www.clubedohardware.com.br",
		IpAddress: "104.20.22.185",
	},
	&fronted.Masquerade{
		Domain:    "www.connectify.me",
		IpAddress: "141.101.113.63",
	},
	&fronted.Masquerade{
		Domain:    "www.cpalead.com",
		IpAddress: "198.41.185.57",
	},
	&fronted.Masquerade{
		Domain:    "www.cryptocoinsnews.com",
		IpAddress: "141.101.124.36",
	},
	&fronted.Masquerade{
		Domain:    "www.cssauthor.com",
		IpAddress: "108.162.206.9",
	},
	&fronted.Masquerade{
		Domain:    "www.cyanogenmod.org",
		IpAddress: "162.159.244.104",
	},
	&fronted.Masquerade{
		Domain:    "www.davidicke.com",
		IpAddress: "198.41.188.87",
	},
	&fronted.Masquerade{
		Domain:    "www.daz3d.com",
		IpAddress: "190.93.241.173",
	},
	&fronted.Masquerade{
		Domain:    "www.desi-tashan.com",
		IpAddress: "162.159.252.243",
	},
	&fronted.Masquerade{
		Domain:    "www.digitalpoint.com",
		IpAddress: "162.159.243.121",
	},
	&fronted.Masquerade{
		Domain:    "www.ekino.tv",
		IpAddress: "162.159.247.209",
	},
	&fronted.Masquerade{
		Domain:    "www.elplural.com",
		IpAddress: "162.159.245.126",
	},
	&fronted.Masquerade{
		Domain:    "www.emailmeform.com",
		IpAddress: "104.16.14.9",
	},
	&fronted.Masquerade{
		Domain:    "www.erepublik.com",
		IpAddress: "198.41.191.76",
	},
	&fronted.Masquerade{
		Domain:    "www.ezilon.com",
		IpAddress: "141.101.112.66",
	},
	&fronted.Masquerade{
		Domain:    "www.fatosdesconhecidos.com.br",
		IpAddress: "198.41.207.179",
	},
	&fronted.Masquerade{
		Domain:    "www.foodpanda.in",
		IpAddress: "104.16.1.10",
	},
	&fronted.Masquerade{
		Domain:    "www.forosdelweb.com",
		IpAddress: "141.101.121.38",
	},
	&fronted.Masquerade{
		Domain:    "www.freeonlinegames.com",
		IpAddress: "190.93.243.38",
	},
	&fronted.Masquerade{
		Domain:    "www.furaffinity.net",
		IpAddress: "104.20.5.196",
	},
	&fronted.Masquerade{
		Domain:    "www.geenstijl.nl",
		IpAddress: "162.159.252.153",
	},
	&fronted.Masquerade{
		Domain:    "www.giltcity.com",
		IpAddress: "141.101.115.238",
	},
	&fronted.Masquerade{
		Domain:    "www.globallshare.com",
		IpAddress: "108.162.200.227",
	},
	&fronted.Masquerade{
		Domain:    "www.grandbux.net",
		IpAddress: "108.162.200.162",
	},
	&fronted.Masquerade{
		Domain:    "www.gulli.com",
		IpAddress: "190.93.241.27",
	},
	&fronted.Masquerade{
		Domain:    "www.hawkhost.com",
		IpAddress: "104.20.0.97",
	},
	&fronted.Masquerade{
		Domain:    "www.hespress.com",
		IpAddress: "162.159.255.97",
	},
	&fronted.Masquerade{
		Domain:    "www.huaweidevice.co.in",
		IpAddress: "198.41.205.132",
	},
	&fronted.Masquerade{
		Domain:    "www.iab.net",
		IpAddress: "141.101.112.75",
	},
	&fronted.Masquerade{
		Domain:    "www.india-forums.com",
		IpAddress: "72.18.156.142",
	},
	&fronted.Masquerade{
		Domain:    "www.infusionsoft.com",
		IpAddress: "198.41.247.139",
	},
	&fronted.Masquerade{
		Domain:    "www.iol.co.za",
		IpAddress: "104.20.29.75",
	},
	&fronted.Masquerade{
		Domain:    "www.jobscore.com",
		IpAddress: "190.93.243.223",
	},
	&fronted.Masquerade{
		Domain:    "www.joe.ie",
		IpAddress: "108.162.202.217",
	},
	&fronted.Masquerade{
		Domain:    "www.jonloomer.com",
		IpAddress: "141.101.125.76",
	},
	&fronted.Masquerade{
		Domain:    "www.joomshaper.com",
		IpAddress: "108.162.206.40",
	},
	&fronted.Masquerade{
		Domain:    "www.jotform.com",
		IpAddress: "141.101.121.43",
	},
	&fronted.Masquerade{
		Domain:    "www.jumia.com.eg",
		IpAddress: "198.41.184.224",
	},
	&fronted.Masquerade{
		Domain:    "www.lebanese-forces.com",
		IpAddress: "141.101.121.65",
	},
	&fronted.Masquerade{
		Domain:    "www.levelup.com",
		IpAddress: "162.159.253.191",
	},
	&fronted.Masquerade{
		Domain:    "www.life.com.tw",
		IpAddress: "141.101.113.19",
	},
	&fronted.Masquerade{
		Domain:    "www.like4like.org",
		IpAddress: "141.101.112.76",
	},
	&fronted.Masquerade{
		Domain:    "www.maduradas.com",
		IpAddress: "162.159.243.224",
	},
	&fronted.Masquerade{
		Domain:    "www.mafa.com",
		IpAddress: "162.159.254.249",
	},
	&fronted.Masquerade{
		Domain:    "www.maxmind.com",
		IpAddress: "141.101.115.190",
	},
	&fronted.Masquerade{
		Domain:    "www.mindtools.com",
		IpAddress: "162.159.252.124",
	},
	&fronted.Masquerade{
		Domain:    "www.mistreci.com",
		IpAddress: "141.101.126.48",
	},
	&fronted.Masquerade{
		Domain:    "www.mkyong.com",
		IpAddress: "108.162.203.6",
	},
	&fronted.Masquerade{
		Domain:    "www.mobofree.com",
		IpAddress: "162.159.255.219",
	},
	&fronted.Masquerade{
		Domain:    "www.modernghana.com",
		IpAddress: "162.159.252.105",
	},
	&fronted.Masquerade{
		Domain:    "www.mp3xd.com",
		IpAddress: "108.162.205.143",
	},
	&fronted.Masquerade{
		Domain:    "www.myvidster.com",
		IpAddress: "198.41.205.6",
	},
	&fronted.Masquerade{
		Domain:    "www.naointendo.com.br",
		IpAddress: "162.159.244.65",
	},
	&fronted.Masquerade{
		Domain:    "www.newgrounds.com",
		IpAddress: "198.41.188.234",
	},
	&fronted.Masquerade{
		Domain:    "www.nthwall.com",
		IpAddress: "104.20.2.28",
	},
	&fronted.Masquerade{
		Domain:    "www.oboom.com",
		IpAddress: "104.20.5.7",
	},
	&fronted.Masquerade{
		Domain:    "www.ofreegames.com",
		IpAddress: "198.41.247.249",
	},
	&fronted.Masquerade{
		Domain:    "www.okcupid.com",
		IpAddress: "198.41.208.133",
	},
	&fronted.Masquerade{
		Domain:    "www.pccomponentes.com",
		IpAddress: "162.159.255.66",
	},
	&fronted.Masquerade{
		Domain:    "www.plugrush.com",
		IpAddress: "162.159.255.156",
	},
	&fronted.Masquerade{
		Domain:    "www.primewire.ag",
		IpAddress: "104.20.31.76",
	},
	&fronted.Masquerade{
		Domain:    "www.problogger.net",
		IpAddress: "162.159.248.46",
	},
	&fronted.Masquerade{
		Domain:    "www.producthunt.com",
		IpAddress: "198.41.207.194",
	},
	&fronted.Masquerade{
		Domain:    "www.pushbullet.com",
		IpAddress: "162.159.242.182",
	},
	&fronted.Masquerade{
		Domain:    "www.quadratin.com.mx",
		IpAddress: "162.159.254.44",
	},
	&fronted.Masquerade{
		Domain:    "www.racing-games.com",
		IpAddress: "162.159.253.249",
	},
	&fronted.Masquerade{
		Domain:    "www.rapidvideo.org",
		IpAddress: "162.159.241.68",
	},
	&fronted.Masquerade{
		Domain:    "www.ratemds.com",
		IpAddress: "104.20.20.13",
	},
	&fronted.Masquerade{
		Domain:    "www.renuevodeplenitud.com",
		IpAddress: "162.159.240.79",
	},
	&fronted.Masquerade{
		Domain:    "www.rome2rio.com",
		IpAddress: "108.162.205.115",
	},
	&fronted.Masquerade{
		Domain:    "www.shortlist.com",
		IpAddress: "190.93.242.31",
	},
	&fronted.Masquerade{
		Domain:    "www.sm3na.com",
		IpAddress: "198.41.179.172",
	},
	&fronted.Masquerade{
		Domain:    "www.somethingawful.com",
		IpAddress: "198.41.187.131",
	},
	&fronted.Masquerade{
		Domain:    "www.ssense.com",
		IpAddress: "104.20.12.4",
	},
	&fronted.Masquerade{
		Domain:    "www.stoiximan.gr",
		IpAddress: "141.101.123.131",
	},
	&fronted.Masquerade{
		Domain:    "www.sundayworld.com",
		IpAddress: "198.41.187.49",
	},
	&fronted.Masquerade{
		Domain:    "www.surveygizmo.com",
		IpAddress: "104.16.19.4",
	},
	&fronted.Masquerade{
		Domain:    "www.sweetfunnycool.com",
		IpAddress: "198.41.249.82",
	},
	&fronted.Masquerade{
		Domain:    "www.techdirt.com",
		IpAddress: "162.159.244.199",
	},
	&fronted.Masquerade{
		Domain:    "www.teefury.com",
		IpAddress: "190.93.243.11",
	},
	&fronted.Masquerade{
		Domain:    "www.thaqafnafsak.com",
		IpAddress: "198.41.207.148",
	},
	&fronted.Masquerade{
		Domain:    "www.thegrommet.com",
		IpAddress: "198.41.188.212",
	},
	&fronted.Masquerade{
		Domain:    "www.theladbible.com",
		IpAddress: "198.41.215.4",
	},
	&fronted.Masquerade{
		Domain:    "www.torrentfunk.com",
		IpAddress: "198.41.200.42",
	},
	&fronted.Masquerade{
		Domain:    "www.traidnt.net",
		IpAddress: "141.101.123.65",
	},
	&fronted.Masquerade{
		Domain:    "www.tvrage.com",
		IpAddress: "141.101.125.12",
	},
	&fronted.Masquerade{
		Domain:    "www.twickerz.com",
		IpAddress: "198.41.249.234",
	},
	&fronted.Masquerade{
		Domain:    "www.vavel.com",
		IpAddress: "190.93.243.103",
	},
	&fronted.Masquerade{
		Domain:    "www.wayn.com",
		IpAddress: "190.93.243.109",
	},
	&fronted.Masquerade{
		Domain:    "www.webmastersitesi.com",
		IpAddress: "141.101.120.108",
	},
	&fronted.Masquerade{
		Domain:    "www.whatismyip.com",
		IpAddress: "141.101.120.14",
	},
	&fronted.Masquerade{
		Domain:    "www.whmcs.com",
		IpAddress: "104.20.20.8",
	},
	&fronted.Masquerade{
		Domain:    "www.zaman.com.tr",
		IpAddress: "190.93.245.170",
	},
	&fronted.Masquerade{
		Domain:    "www.zopim.com",
		IpAddress: "190.93.240.200",
	},
	&fronted.Masquerade{
		Domain:    "www.zumba.com",
		IpAddress: "190.93.247.77",
	},
	&fronted.Masquerade{
		Domain:    "x-kom.pl",
		IpAddress: "104.20.29.24",
	},
	&fronted.Masquerade{
		Domain:    "xat.com",
		IpAddress: "141.101.113.82",
	},
	&fronted.Masquerade{
		Domain:    "yifysubtitles.com",
		IpAddress: "141.101.127.79",
	},
	&fronted.Masquerade{
		Domain:    "youm7.com",
		IpAddress: "104.16.16.116",
	},
	&fronted.Masquerade{
		Domain:    "yourbittorrent.com",
		IpAddress: "198.41.203.40",
	},
	&fronted.Masquerade{
		Domain:    "yourdailyscoop.com",
		IpAddress: "162.159.249.210",
	},
	&fronted.Masquerade{
		Domain:    "yyv.co",
		IpAddress: "190.93.240.65",
	},
	&fronted.Masquerade{
		Domain:    "z6.com",
		IpAddress: "162.159.247.121",
	},
	&fronted.Masquerade{
		Domain:    "zennolab.com",
		IpAddress: "144.76.44.168",
	},
	&fronted.Masquerade{
		Domain:    "zentrum-der-gesundheit.de",
		IpAddress: "141.101.113.102",
	},
	&fronted.Masquerade{
		Domain:    "zerozero.pt",
		IpAddress: "198.41.190.107",
	},
	&fronted.Masquerade{
		Domain:    "zurb.com",
		IpAddress: "104.20.4.2",
	},
	&fronted.Masquerade{
		Domain:    "zwaar.net",
		IpAddress: "162.159.247.231",
	},
}
