package main

import (
	"fmt"
	"github.com/kgori/go_phylo/lexer"
)

func main() {
	l := lexer.New("test", "(Equus_caballus:0.0657603[&&NHX:name=Equus_caballus],(Manis_pentadactyla:0.228752[&&NHX:name=Manis_pentadactyla],((Nandinia_binotata:0.0824088[&&NHX:name=Nandinia_binotata],(((Prionodon_pardicolor:0.0119788[&&NHX:name=Prionodon_pardicolor],Prionodon_linsang:0.0370715[&&NHX:name=Prionodon_linsang])1:0.0542796[&&NHX:name=15],((Homotherium_serum:0.0463723[&&NHX:name=Homotherium_serum],Smilodon_populator:0.0235564[&&NHX:name=Smilodon_populator])1:0.0109816[&&NHX:name=25],(((Neofelis_nebulosa_diardi:0.00625375[&&NHX:name=Neofelis_nebulosa_diardi],Neofelis_nebulosa:0.00632091[&&NHX:name=Neofelis_nebulosa])1:0.0125728[&&NHX:name=64],(Panthera_tigris:0.0141581[&&NHX:name=Panthera_tigris],(Uncia_uncia:0.0120819[&&NHX:name=Uncia_uncia],(Panthera_onca:0.00942411[&&NHX:name=Panthera_onca],(Panthera_pardus:0.00892205[&&NHX:name=Panthera_pardus],(Panthera_leo_atrox:0.00598901[&&NHX:name=Panthera_leo_atrox],Panthera_leo:0.00298375[&&NHX:name=Panthera_leo])1:0.00424102[&&NHX:name=173])1:0.00159918[&&NHX:name=146])1:0.00110658[&&NHX:name=118])1:0.00330354[&&NHX:name=92])1:0.00550428[&&NHX:name=65])1:0.00821594[&&NHX:name=42],((Pardofelis_marmorata:0.0121592[&&NHX:name=Pardofelis_marmorata],(Catopuma_temminckii:0.00974119[&&NHX:name=Catopuma_temminckii],Catopuma_badia:0.00648451[&&NHX:name=Catopuma_badia])1:0.0031053[&&NHX:name=93])1:0.00513623[&&NHX:name=66],((Leptailurus_serval:0.0171149[&&NHX:name=Leptailurus_serval],(Caracal_caracal:0.004777[&&NHX:name=Caracal_caracal],Profelis_aurata:0.00597792[&&NHX:name=Profelis_aurata])1:0.00739043[&&NHX:name=119])1:0.00520044[&&NHX:name=94],((Leopardus_pardalis:0.00717691[&&NHX:name=Leopardus_pardalis],((Leopardus_wiedii:0.00501133[&&NHX:name=Leopardus_wiedii],Leopardus_jacobitus:0.00636775[&&NHX:name=Leopardus_jacobitus])1:0.00121921[&&NHX:name=174],(Leopardus_colocolo:0.00468571[&&NHX:name=Leopardus_colocolo],(Leopardus_tigrinus:0.00502121[&&NHX:name=Leopardus_tigrinus],(Leopardus_geoffroyi:0.00237615[&&NHX:name=Leopardus_geoffroyi],Leopardus_guigna:0.00289117[&&NHX:name=Leopardus_guigna])1:0.00277333[&&NHX:name=214])1:0.00140522[&&NHX:name=197])1:0.00222766[&&NHX:name=175])1:0.00207789[&&NHX:name=147])1:0.00996137[&&NHX:name=120],(((Lynx_rufus:0.0102263[&&NHX:name=Lynx_rufus],(Lynx_pardinus:0.0037651[&&NHX:name=Lynx_pardinus],(Lynx_lynx:0.00313157[&&NHX:name=Lynx_lynx],Lynx_canadensis:0.00795745[&&NHX:name=Lynx_canadensis])1:0.00129943[&&NHX:name=215])1:0.00460589[&&NHX:name=198])1:0.00629346[&&NHX:name=176],(Acinonyx_jubatus:0.0149736[&&NHX:name=Acinonyx_jubatus],(Puma_yagouaroundi:0.0106215[&&NHX:name=Puma_yagouaroundi],(Miracinonyx_trumani:0.0118041[&&NHX:name=Miracinonyx_trumani],Puma_concolor:0.0051867[&&NHX:name=Puma_concolor])1:0.00573632[&&NHX:name=216])1:0.00206854[&&NHX:name=199])1:0.00398086[&&NHX:name=177])1:0.00157056[&&NHX:name=148],(Felis_manul:0.0162714[&&NHX:name=Felis_manul],((Prionailurus_rubiginosa:0.00937969[&&NHX:name=Prionailurus_rubiginosa],(Prionailurus_planiceps:0.00836189[&&NHX:name=Prionailurus_planiceps],(Prionailurus_iriomotensis:0.00617196[&&NHX:name=Prionailurus_iriomotensis],(Prionailurus_viverrinus:0.00662689[&&NHX:name=Prionailurus_viverrinus],Prionailurus_bengalensis:0.00547955[&&NHX:name=Prionailurus_bengalensis])1:0.000684772[&&NHX:name=235])1:0.00132553[&&NHX:name=226])1:0.00507222[&&NHX:name=217])1:0.00329225[&&NHX:name=200],(Felis_chaus:0.00729248[&&NHX:name=Felis_chaus],(Felis_nigripes:0.00669501[&&NHX:name=Felis_nigripes],(Felis_margarita:0.0043037[&&NHX:name=Felis_margarita],(Felis_bieti:0.00196434[&&NHX:name=Felis_bieti],(Felis_silvestris:0.000751818[&&NHX:name=Felis_silvestris],Felis_catus:0.00236102[&&NHX:name=Felis_catus])1:0.00100488[&&NHX:name=246])1:0.00159317[&&NHX:name=236])1:0.00334642[&&NHX:name=227])1:0.0011093[&&NHX:name=218])1:0.00764391[&&NHX:name=201])1:0.00103837[&&NHX:name=178])1:0.0016953[&&NHX:name=149])1:0.00140862[&&NHX:name=121])1:0.000531061[&&NHX:name=95])1:0.00108552[&&NHX:name=67])1:0.00304258[&&NHX:name=43])1:0.0142919[&&NHX:name=26])1:0.0271661[&&NHX:name=16])1:0.00895768[&&NHX:name=9],((((Cynogale_bennettii:0.0474279[&&NHX:name=Cynogale_bennettii],(Chrotogale_owstoni:0.0222648[&&NHX:name=Chrotogale_owstoni],Hemigalus_derbyanus:0.0245623[&&NHX:name=Hemigalus_derbyanus])1:0.00985628[&&NHX:name=68])1:0.0168038[&&NHX:name=44],(Arctogalidia_trivirgata:0.0329623[&&NHX:name=Arctogalidia_trivirgata],(Arctictis_binturong:0.0267927[&&NHX:name=Arctictis_binturong],(Paguma_larvata:0.0171751[&&NHX:name=Paguma_larvata],(Paradoxurus_hermaphroditus:0.0233677[&&NHX:name=Paradoxurus_hermaphroditus],(Paradoxurus_jerdoni:0.00794813[&&NHX:name=Paradoxurus_jerdoni],Paradoxurus_zeylonensis:0.0038707[&&NHX:name=Paradoxurus_zeylonensis])1:0.00625437[&&NHX:name=150])1:0.0108124[&&NHX:name=122])1:0.0144129[&&NHX:name=96])1:0.0116034[&&NHX:name=69])1:0.00518012[&&NHX:name=45])1:0.00945874[&&NHX:name=27],((Viverricula_indica:0.0217108[&&NHX:name=Viverricula_indica],(Civettictis_civetta:0.0262972[&&NHX:name=Civettictis_civetta],(Viverra_tangalunga:0.00836972[&&NHX:name=Viverra_tangalunga],(Viverra_megaspila:0.00728639[&&NHX:name=Viverra_megaspila],Viverra_zibetha:0.00556808[&&NHX:name=Viverra_zibetha])1:0.00324208[&&NHX:name=123])1:0.0105118[&&NHX:name=97])1:0.00755861[&&NHX:name=70])1:0.0102854[&&NHX:name=46],(Poiana_richardsonii:0.015198[&&NHX:name=Poiana_richardsonii],(Genetta_thierryi:0.0142067[&&NHX:name=Genetta_thierryi],(Genetta_victoriae:0.0022378[&&NHX:name=Genetta_victoriae],((Genetta_piscivora:0.0070042[&&NHX:name=Genetta_piscivora],(Genetta_johnstoni:0.00882213[&&NHX:name=Genetta_johnstoni],Genetta_genetta_felina:0.0399525[&&NHX:name=Genetta_genetta_felina])1:0.000428873[&&NHX:name=179])1:0.00199252[&&NHX:name=151],((Genetta_cristata:0.00302717[&&NHX:name=Genetta_cristata],Genetta_servalina:0.00332813[&&NHX:name=Genetta_servalina])1:0.00416751[&&NHX:name=180],(Genetta_genetta:0.00991708[&&NHX:name=Genetta_genetta],(Genetta_angolensis:0.00388506[&&NHX:name=Genetta_angolensis],((Genetta_bourloni:0.00154388[&&NHX:name=Genetta_bourloni],(Genetta_poensis:0.00107338[&&NHX:name=Genetta_poensis],Genetta_pardina:0.00470004[&&NHX:name=Genetta_pardina])1:0.000646184[&&NHX:name=237])1:0.00252769[&&NHX:name=228],(Genetta_maculata:0.00085734[&&NHX:name=Genetta_maculata],(Genetta_tigrina:0.00207606[&&NHX:name=Genetta_tigrina],Genetta_rubiginosa:2.98487e-06[&&NHX:name=Genetta_rubiginosa])1:0.000809844[&&NHX:name=238])1:0.00183875[&&NHX:name=229])1:0.000727704[&&NHX:name=219])1:0.0029901[&&NHX:name=202])1:0.00200309[&&NHX:name=181])1:0.000913183[&&NHX:name=152])1:0.00375073[&&NHX:name=124])1:0.00185189[&&NHX:name=98])1:0.00927256[&&NHX:name=71])1:0.0178316[&&NHX:name=47])1:0.00792004[&&NHX:name=28])1:0.0201444[&&NHX:name=17],((Proteles_cristata:0.0193368[&&NHX:name=Proteles_cristata],(Crocuta_crocuta:0.0115486[&&NHX:name=Crocuta_crocuta],(Hyaena_hyaena:0.00775834[&&NHX:name=Hyaena_hyaena],Hyaena_brunnea:0.00642005[&&NHX:name=Hyaena_brunnea])1:0.00794892[&&NHX:name=72])1:0.00331739[&&NHX:name=48])1:0.056573[&&NHX:name=29],(((Eupleres_goudotii:0.0356391[&&NHX:name=Eupleres_goudotii],Fossa_fossana:0.0157361[&&NHX:name=Fossa_fossana])1:0.0173803[&&NHX:name=73],(Cryptoprocta_ferox:0.0334541[&&NHX:name=Cryptoprocta_ferox],(Galidia_elegans:0.0204042[&&NHX:name=Galidia_elegans],(Galidictis_fasciata:0.0101187[&&NHX:name=Galidictis_fasciata],(Salanoia_concolor:0.00540092[&&NHX:name=Salanoia_concolor],Mungotictis_decemlineata:0.00637119[&&NHX:name=Mungotictis_decemlineata])1:0.00573774[&&NHX:name=153])1:0.00817788[&&NHX:name=125])1:0.0130302[&&NHX:name=99])1:0.00213336[&&NHX:name=74])1:0.00882353[&&NHX:name=49],((Suricata_suricatta:0.0356739[&&NHX:name=Suricata_suricatta],((Mungos_mungo:0.0238891[&&NHX:name=Mungos_mungo],Liberiictis_kuhni:0.020454[&&NHX:name=Liberiictis_kuhni])1:0.00187775[&&NHX:name=126],((Crossarchus_alexandri:0.00891815[&&NHX:name=Crossarchus_alexandri],Crossarchus_obscurus:0.00843026[&&NHX:name=Crossarchus_obscurus])1:0.009001[&&NHX:name=154],(Helogale_parvula:0.0108221[&&NHX:name=Helogale_parvula],Helogale_hirtula:0.0148225[&&NHX:name=Helogale_hirtula])1:0.00632769[&&NHX:name=155])1:0.0060126[&&NHX:name=127])1:0.00190396[&&NHX:name=100])1:0.00546444[&&NHX:name=75],(((Herpestes_ichneumon:0.013146[&&NHX:name=Herpestes_ichneumon],(Galerella_pulverulenta:0.0140981[&&NHX:name=Galerella_pulverulenta],Galerella_sanguinea:0.0163908[&&NHX:name=Galerella_sanguinea])1:0.00134554[&&NHX:name=156])1:0.0033822[&&NHX:name=128],(Ichneumia_albicauda:0.0163631[&&NHX:name=Ichneumia_albicauda],((Paracynictis_selousi:0.00601464[&&NHX:name=Paracynictis_selousi],Cynictis_penicillata:0.00926832[&&NHX:name=Cynictis_penicillata])1:0.00531351[&&NHX:name=182],(Rhynchogale_melleri:0.0147368[&&NHX:name=Rhynchogale_melleri],(Bdeogale_crassicauda:0.00869779[&&NHX:name=Bdeogale_crassicauda],Bdeogale_nigripes:0.00593805[&&NHX:name=Bdeogale_nigripes])1:0.0101432[&&NHX:name=203])1:0.00130513[&&NHX:name=183])1:0.00263635[&&NHX:name=157])1:0.00435658[&&NHX:name=129])1:0.00226884[&&NHX:name=101],((Atilax_paludinosus:0.0123701[&&NHX:name=Atilax_paludinosus],Herpestes_naso:0.0164471[&&NHX:name=Herpestes_naso])1:0.00461777[&&NHX:name=130],((Herpestes_urva:0.0144381[&&NHX:name=Herpestes_urva],(Herpestes_smithii:0.0137231[&&NHX:name=Herpestes_smithii],Herpestes_brachyurus:0.0110097[&&NHX:name=Herpestes_brachyurus])1:0.0016185[&&NHX:name=184])1:0.0038491[&&NHX:name=158],(Herpestes_fuscus:0.0160018[&&NHX:name=Herpestes_fuscus],(Herpestes_edwardsii:0.00907864[&&NHX:name=Herpestes_edwardsii],(Herpestes_javanicus:0.00442893[&&NHX:name=Herpestes_javanicus],Herpestes_javanicus_auropunctatus:0.00528129[&&NHX:name=Herpestes_javanicus_auropunctatus])1:0.00213985[&&NHX:name=204])1:0.00477914[&&NHX:name=185])1:0.00457005[&&NHX:name=159])1:0.00147699[&&NHX:name=131])1:0.000879197[&&NHX:name=102])1:0.00699667[&&NHX:name=76])1:0.0378648[&&NHX:name=50])1:0.0113685[&&NHX:name=30])1:0.0148876[&&NHX:name=18])1:0.00461059[&&NHX:name=10])1:0.0143909[&&NHX:name=6])1:0.0377696[&&NHX:name=4],(((Nyctereutes_procyonoides:0.0318632[&&NHX:name=Nyctereutes_procyonoides],(Otocyon_megalotis:0.0290466[&&NHX:name=Otocyon_megalotis],((Urocyon_cinereoargenteus:0.000644727[&&NHX:name=Urocyon_cinereoargenteus],Urocyon_littoralis:0.00272301[&&NHX:name=Urocyon_littoralis])1:0.0231217[&&NHX:name=51],((Vulpes_zerda:0.00919051[&&NHX:name=Vulpes_zerda],Vulpes_cana:0.00717477[&&NHX:name=Vulpes_cana])1:0.00322632[&&NHX:name=77],(Vulpes_chama:0.00959538[&&NHX:name=Vulpes_chama],((Vulpes_lagopus:0.00480341[&&NHX:name=Vulpes_lagopus],(Vulpes_velox:0.00106071[&&NHX:name=Vulpes_velox],Vulpes_macrotis:2.98487e-06[&&NHX:name=Vulpes_macrotis])1:0.00267026[&&NHX:name=160])1:0.00882313[&&NHX:name=132],(Vulpes_corsac:0.00650357[&&NHX:name=Vulpes_corsac],(Vulpes_ferrilata:0.00987012[&&NHX:name=Vulpes_ferrilata],(Vulpes_rueppellii:0.00409731[&&NHX:name=Vulpes_rueppellii],Vulpes_vulpes:0.00699519[&&NHX:name=Vulpes_vulpes])1:0.00134785[&&NHX:name=186])1:0.000887467[&&NHX:name=161])1:0.00318424[&&NHX:name=133])1:0.00363995[&&NHX:name=103])1:0.00156919[&&NHX:name=78])1:0.00926249[&&NHX:name=52])1:0.00246293[&&NHX:name=31])1:0.00207356[&&NHX:name=19])1:0.00349277[&&NHX:name=11],(((Dusicyon_australis:0.0247355[&&NHX:name=Dusicyon_australis],(Speothos_venaticus:0.0122515[&&NHX:name=Speothos_venaticus],Chrysocyon_brachyurus:0.0123845[&&NHX:name=Chrysocyon_brachyurus])1:0.00312936[&&NHX:name=53])1:0.00141821[&&NHX:name=32],((Cerdocyon_thous:0.00691267[&&NHX:name=Cerdocyon_thous],Atelocynus_microtis:0.00631715[&&NHX:name=Atelocynus_microtis])1:0.00103032[&&NHX:name=54],(Lycalopex_sechurae:0.00376222[&&NHX:name=Lycalopex_sechurae],(Lycalopex_gymnocercus:0.00207548[&&NHX:name=Lycalopex_gymnocercus],((Lycalopex_fulvipes:0.00390457[&&NHX:name=Lycalopex_fulvipes],Lycalopex_vetulus:0.00400671[&&NHX:name=Lycalopex_vetulus])1:0.00089609[&&NHX:name=134],(Lycalopex_culpaeus:0.00211328[&&NHX:name=Lycalopex_culpaeus],Lycalopex_griseus:0.0016514[&&NHX:name=Lycalopex_griseus])1:0.000436069[&&NHX:name=135])1:0.000284261[&&NHX:name=104])1:0.00110526[&&NHX:name=79])1:0.00357853[&&NHX:name=55])1:0.00496381[&&NHX:name=33])1:0.00237717[&&NHX:name=20],((Canis_adustus:0.00757918[&&NHX:name=Canis_adustus],Canis_mesomelas:0.00835711[&&NHX:name=Canis_mesomelas])1:0.0032062[&&NHX:name=34],(Lycaon_pictus:0.0150674[&&NHX:name=Lycaon_pictus],((Canis_lupus_rufus:0.000625937[&&NHX:name=Canis_lupus_rufus],Cuon_alpinus:0.0112423[&&NHX:name=Cuon_alpinus])1:0.00653274[&&NHX:name=80],(Canis_simensis:0.00558846[&&NHX:name=Canis_simensis],((Canis_lupus_lycaon:0.00117201[&&NHX:name=Canis_lupus_lycaon],Canis_latrans:0.000251599[&&NHX:name=Canis_latrans])1:0.00585974[&&NHX:name=136],(Canis_aureus:0.00531187[&&NHX:name=Canis_aureus],(Canis_lupus_laniger:0.00611223[&&NHX:name=Canis_lupus_laniger],(Canis_lupus_pallipes:0.00202497[&&NHX:name=Canis_lupus_pallipes],(Canis_lupus_familiaris:0.00410876[&&NHX:name=Canis_lupus_familiaris],Canis_lupus:0.000674282[&&NHX:name=Canis_lupus])1:0.000997299[&&NHX:name=205])1:0.00129527[&&NHX:name=187])1:0.0018406[&&NHX:name=162])1:0.00104607[&&NHX:name=137])1:0.00180712[&&NHX:name=105])1:0.0027204[&&NHX:name=81])1:0.00228441[&&NHX:name=56])1:0.00126876[&&NHX:name=35])1:0.00247272[&&NHX:name=21])1:0.0104682[&&NHX:name=12])1:0.112063[&&NHX:name=7],((Ailuropoda_melanoleuca:0.0384281[&&NHX:name=Ailuropoda_melanoleuca],((Arctodus_simus:0.0160215[&&NHX:name=Arctodus_simus],Tremarctos_ornatus:0.0156429[&&NHX:name=Tremarctos_ornatus])1:0.012709[&&NHX:name=36],(Melursus_ursinus:0.0134346[&&NHX:name=Melursus_ursinus],((Ursus_spelaeus:0.00946898[&&NHX:name=Ursus_spelaeus],(Ursus_maritimus:0.00255419[&&NHX:name=Ursus_maritimus],Ursus_arctos:0.00262755[&&NHX:name=Ursus_arctos])1:0.00400774[&&NHX:name=106])1:0.00546319[&&NHX:name=82],(Helarctos_malayanus:0.0130452[&&NHX:name=Helarctos_malayanus],(Ursus_thibetanus:0.00971481[&&NHX:name=Ursus_thibetanus],Ursus_americanus:0.012135[&&NHX:name=Ursus_americanus])1:0.00141312[&&NHX:name=107])1:0.00126094[&&NHX:name=83])1:0.00196492[&&NHX:name=57])1:0.0163426[&&NHX:name=37])1:0.00733988[&&NHX:name=22])1:0.0521379[&&NHX:name=13],(((Odobenus_rosmarus:0.0314261[&&NHX:name=Odobenus_rosmarus],(Callorhinus_ursinus:0.0152407[&&NHX:name=Callorhinus_ursinus],((Eumetopias_jubatus:0.0092917[&&NHX:name=Eumetopias_jubatus],(Zalophus_wollebaeki:0.000505899[&&NHX:name=Zalophus_wollebaeki],Zalophus_californianus:2.98487e-06[&&NHX:name=Zalophus_californianus])1:0.00693574[&&NHX:name=138])1:0.00168582[&&NHX:name=108],(Arctocephalus_pusillus:0.011504[&&NHX:name=Arctocephalus_pusillus],(Otaria_flavescens:0.0134479[&&NHX:name=Otaria_flavescens],((Neophoca_cinerea:0.0106044[&&NHX:name=Neophoca_cinerea],Phocarctos_hookeri:0.0107091[&&NHX:name=Phocarctos_hookeri])1:0.000941293[&&NHX:name=188],((Arctocephalus_galapagoensis:0.000763006[&&NHX:name=Arctocephalus_galapagoensis],(Arctocephalus_australis:0.00232508[&&NHX:name=Arctocephalus_australis],Arctocephalus_forsteri:0.003046[&&NHX:name=Arctocephalus_forsteri])1:0.00297236[&&NHX:name=220])1:0.00198312[&&NHX:name=206],((Arctocephalus_townsendi:0.000555156[&&NHX:name=Arctocephalus_townsendi],Arctocephalus_philippii:0.00147929[&&NHX:name=Arctocephalus_philippii])1:0.0073355[&&NHX:name=221],(Arctocephalus_gazella:0.00129432[&&NHX:name=Arctocephalus_gazella],Arctocephalus_tropicalis:0.00039568[&&NHX:name=Arctocephalus_tropicalis])1:0.00520575[&&NHX:name=222])1:0.000466117[&&NHX:name=207])1:0.00194597[&&NHX:name=189])1:0.000972263[&&NHX:name=163])1:0.000606503[&&NHX:name=139])1:0.00309974[&&NHX:name=109])1:0.00586133[&&NHX:name=84])1:0.0151206[&&NHX:name=58])1:0.0177416[&&NHX:name=38],(((Monachus_schauinslandi:0.0191814[&&NHX:name=Monachus_schauinslandi],Monachus_monachus:0.0195133[&&NHX:name=Monachus_monachus])1:0.00326035[&&NHX:name=85],((Mirounga_angustirostris:0.00551802[&&NHX:name=Mirounga_angustirostris],Mirounga_leonina:0.00446953[&&NHX:name=Mirounga_leonina])1:0.0156227[&&NHX:name=110],(Lobodon_carcinophaga:0.0111364[&&NHX:name=Lobodon_carcinophaga],(Ommatophoca_rossii:0.0130442[&&NHX:name=Ommatophoca_rossii],(Hydrurga_leptonyx:0.00837358[&&NHX:name=Hydrurga_leptonyx],Leptonychotes_weddellii:0.00721588[&&NHX:name=Leptonychotes_weddellii])1:0.00345218[&&NHX:name=164])1:0.00120789[&&NHX:name=140])1:0.00329933[&&NHX:name=111])1:0.00298488[&&NHX:name=86])1:0.007712[&&NHX:name=59],(Erignathus_barbatus:0.0199005[&&NHX:name=Erignathus_barbatus],(Cystophora_cristata:0.0151178[&&NHX:name=Cystophora_cristata],((Histriophoca_fasciata:0.0102936[&&NHX:name=Histriophoca_fasciata],Pagophilus_groenlandicus:0.00894045[&&NHX:name=Pagophilus_groenlandicus])1:0.00347339[&&NHX:name=141],((Phoca_vitulina:0.00252294[&&NHX:name=Phoca_vitulina],Phoca_largha:0.00216667[&&NHX:name=Phoca_largha])1:0.00207836[&&NHX:name=165],((Pusa_hispida:0.00377205[&&NHX:name=Pusa_hispida],Pusa_sibirica:0.00290083[&&NHX:name=Pusa_sibirica])1:0.000256205[&&NHX:name=190],(Halichoerus_grypus:0.00438823[&&NHX:name=Halichoerus_grypus],Pusa_caspica:0.0036339[&&NHX:name=Pusa_caspica])1:0.000351598[&&NHX:name=191])1:0.000858185[&&NHX:name=166])1:0.00732045[&&NHX:name=142])1:0.00312745[&&NHX:name=112])1:0.00729196[&&NHX:name=87])1:0.00387443[&&NHX:name=60])1:0.0117345[&&NHX:name=39])1:0.0299714[&&NHX:name=23],(((Mydaus_marchei:0.0228222[&&NHX:name=Mydaus_marchei],Mydaus_javanensis:0.0149022[&&NHX:name=Mydaus_javanensis])1:0.0605577[&&NHX:name=61],((Conepatus_chinga:0.0123923[&&NHX:name=Conepatus_chinga],(Conepatus_leuconotus:0.00148881[&&NHX:name=Conepatus_leuconotus],Conepatus_leuconotus_leuconotus:2.98487e-06[&&NHX:name=Conepatus_leuconotus_leuconotus])1:0.0122481[&&NHX:name=113])1:0.0285227[&&NHX:name=88],((Mephitis_macroura:0.0215602[&&NHX:name=Mephitis_macroura],Mephitis_mephitis:2.98487e-06[&&NHX:name=Mephitis_mephitis])1:0.0363563[&&NHX:name=114],(Spilogale_gracilis:0.0018729[&&NHX:name=Spilogale_gracilis],Spilogale_putorius:0.0270132[&&NHX:name=Spilogale_putorius])1:0.00496451[&&NHX:name=115])1:0.0149043[&&NHX:name=89])1:0.0284759[&&NHX:name=62])1:0.042818[&&NHX:name=40],(Ailurus_fulgens:0.0783653[&&NHX:name=Ailurus_fulgens],((Potos_flavus:0.0477121[&&NHX:name=Potos_flavus],(((Bassariscus_astutus:0.0221631[&&NHX:name=Bassariscus_astutus],Bassariscus_sumichrasti:0.0201268[&&NHX:name=Bassariscus_sumichrasti])1:0.00586528[&&NHX:name=167],(Procyon_cancrivorus:0.011593[&&NHX:name=Procyon_cancrivorus],(Bassaricyon_beddardi:0.0320635[&&NHX:name=Bassaricyon_beddardi],Procyon_lotor:0.00908873[&&NHX:name=Procyon_lotor])1:2.98487e-06[&&NHX:name=192])1:0.0213462[&&NHX:name=168])1:0.0283971[&&NHX:name=143],((Bassaricyon_gabbii:0.00793393[&&NHX:name=Bassaricyon_gabbii],Bassaricyon_alleni:0.00414709[&&NHX:name=Bassaricyon_alleni])1:0.0272188[&&NHX:name=169],(Nasua_nasua:0.0122124[&&NHX:name=Nasua_nasua],(Nasua_narica:0.00756331[&&NHX:name=Nasua_narica],(Nasuella_olivacea_meridensis:0.0127445[&&NHX:name=Nasuella_olivacea_meridensis],Nasuella_olivacea:0.00995691[&&NHX:name=Nasuella_olivacea])1:0.0106975[&&NHX:name=208])1:0.0050318[&&NHX:name=193])1:0.0220308[&&NHX:name=170])1:0.0221905[&&NHX:name=144])1:0.0122021[&&NHX:name=116])1:0.0177867[&&NHX:name=90],(Taxidea_taxus:0.0421426[&&NHX:name=Taxidea_taxus],(Mellivora_capensis:0.041264[&&NHX:name=Mellivora_capensis],((Arctonyx_collaris:0.0178414[&&NHX:name=Arctonyx_collaris],(Meles_meles:0.00819929[&&NHX:name=Meles_meles],(Meles_leucurus:0.00545473[&&NHX:name=Meles_leucurus],Meles_anakuma:0.00769486[&&NHX:name=Meles_anakuma])1:0.00375468[&&NHX:name=209])1:0.00347828[&&NHX:name=194])1:0.0305977[&&NHX:name=171],(((Eira_barbara:0.0296382[&&NHX:name=Eira_barbara],Martes_pennanti:0.0229882[&&NHX:name=Martes_pennanti])1:0.00119316[&&NHX:name=210],(Gulo_gulo:0.0214418[&&NHX:name=Gulo_gulo],(Martes_flavigula:0.0236727[&&NHX:name=Martes_flavigula],(Martes_foina:0.0111102[&&NHX:name=Martes_foina],((Martes_americana_caurina:2.98487e-06[&&NHX:name=Martes_americana_caurina],Martes_americana:0.000572723[&&NHX:name=Martes_americana])1:0.00676482[&&NHX:name=247],(Martes_melampus:0.00388535[&&NHX:name=Martes_melampus],(Martes_martes:0.00448105[&&NHX:name=Martes_martes],Martes_zibellina:0.00291205[&&NHX:name=Martes_zibellina])1:0.00168638[&&NHX:name=254])1:0.0011001[&&NHX:name=248])1:0.00397211[&&NHX:name=239])1:0.00781412[&&NHX:name=230])1:0.00421501[&&NHX:name=223])1:0.00249095[&&NHX:name=211])1:0.0106287[&&NHX:name=195],((Melogale_personata:0.00763884[&&NHX:name=Melogale_personata],Melogale_moschata:0.00516514[&&NHX:name=Melogale_moschata])1:0.0317801[&&NHX:name=212],(((Neovison_vison:0.0137539[&&NHX:name=Neovison_vison],(Mustela_frenata:0.00740708[&&NHX:name=Mustela_frenata],(Mustela_felipei:0.0101227[&&NHX:name=Mustela_felipei],Mustela_africana:0.0114755[&&NHX:name=Mustela_africana])1:0.00236127[&&NHX:name=249])1:0.00532978[&&NHX:name=240])1:0.0130844[&&NHX:name=231],((Mustela_strigidorsa:0.0164702[&&NHX:name=Mustela_strigidorsa],Mustela_nudipes:0.015536[&&NHX:name=Mustela_nudipes])1:0.0125869[&&NHX:name=241],(Mustela_kathiah:0.025357[&&NHX:name=Mustela_kathiah],(Mustela_erminea:0.0119639[&&NHX:name=Mustela_erminea],((Mustela_altaica:0.0114445[&&NHX:name=Mustela_altaica],Mustela_nivalis:0.00831801[&&NHX:name=Mustela_nivalis])1:0.00290194[&&NHX:name=259],(Mustela_itatsi:0.00940627[&&NHX:name=Mustela_itatsi],(Mustela_sibirica:0.0114065[&&NHX:name=Mustela_sibirica],(Mustela_lutreola:0.00365686[&&NHX:name=Mustela_lutreola],(Mustela_putorius:0.00390957[&&NHX:name=Mustela_putorius],(Mustela_eversmanii:0.0015351[&&NHX:name=Mustela_eversmanii],Mustela_nigripes:0.0041513[&&NHX:name=Mustela_nigripes])1:0.000576422[&&NHX:name=269])1:0.00248659[&&NHX:name=268])1:0.00464633[&&NHX:name=266])1:0.00117125[&&NHX:name=263])1:0.00379903[&&NHX:name=260])1:0.00475236[&&NHX:name=255])1:0.00338624[&&NHX:name=250])1:0.00478622[&&NHX:name=242])1:0.00490636[&&NHX:name=232])1:0.0138587[&&NHX:name=224],(((Galictis_cuja:0.0126711[&&NHX:name=Galictis_cuja],Galictis_vittata:0.0103349[&&NHX:name=Galictis_vittata])1:0.0322175[&&NHX:name=243],(Vormela_peregusna:0.0266751[&&NHX:name=Vormela_peregusna],(Ictonyx_libyca:0.0289804[&&NHX:name=Ictonyx_libyca],(Ictonyx_striatus:0.0152787[&&NHX:name=Ictonyx_striatus],Poecilogale_albinucha:0.0170143[&&NHX:name=Poecilogale_albinucha])1:0.00334725[&&NHX:name=256])1:0.00670355[&&NHX:name=251])1:0.0129302[&&NHX:name=244])1:0.00442234[&&NHX:name=233],(Pteronura_brasiliensis:0.0317287[&&NHX:name=Pteronura_brasiliensis],((Lontra_canadensis:0.0131555[&&NHX:name=Lontra_canadensis],(Lontra_longicaudis:0.0054153[&&NHX:name=Lontra_longicaudis],(Lontra_provocax:0.000580748[&&NHX:name=Lontra_provocax],Lontra_felina:0.00251131[&&NHX:name=Lontra_felina])1:0.00291147[&&NHX:name=261])1:0.00921743[&&NHX:name=257])1:0.0172514[&&NHX:name=252],(Enhydra_lutris:0.0230251[&&NHX:name=Enhydra_lutris],(Hydrictis_maculicollis:0.0193912[&&NHX:name=Hydrictis_maculicollis],((Lutra_lutra:0.00800243[&&NHX:name=Lutra_lutra],Lutra_sumatrana:0.00622748[&&NHX:name=Lutra_sumatrana])1:0.00893104[&&NHX:name=264],(Aonyx_capensis:0.0116055[&&NHX:name=Aonyx_capensis],(Aonyx_cinerea:0.0110604[&&NHX:name=Aonyx_cinerea],Lutrogale_perspicillata:0.00689025[&&NHX:name=Lutrogale_perspicillata])1:0.00864706[&&NHX:name=267])1:0.0044999[&&NHX:name=265])1:0.00590028[&&NHX:name=262])1:0.000886077[&&NHX:name=258])1:0.0057041[&&NHX:name=253])1:0.00297242[&&NHX:name=245])1:0.00376246[&&NHX:name=234])1:0.00157997[&&NHX:name=225])1:0.00574969[&&NHX:name=213])1:0.00187305[&&NHX:name=196])1:0.00246221[&&NHX:name=172])1:0.00164624[&&NHX:name=145])1:0.00849895[&&NHX:name=117])1:0.0378267[&&NHX:name=91])1:0.00834901[&&NHX:name=63])1:0.00502436[&&NHX:name=41])1:0.0187344[&&NHX:name=24])1:0.00504491[&&NHX:name=14])1:0.0232955[&&NHX:name=8])1:0.0225588[&&NHX:name=5])1:0.0530296[&&NHX:name=3])1:0.0657603[&&NHX:name=2])[&&NHX:name=1];", 5)
	t := lexer.NewItem()
	fmt.Printf("%T\n", t)

	for !t.IsEOF() {
		t = l.NextItem()
	}
}

// func lex() {
// 	return nil
// }