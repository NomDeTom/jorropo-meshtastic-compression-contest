# Meshtastic Compression Showdown

This project contains benchmarks of various compression algorithms applied on a dataset of meshtastic packets.

For context a Reciprocal Compression Ratio **above** 1 means the compressed data is **bigger** than the uncompressed data.
A ratio **below** 1 means the compressed data is **smaller** than the uncompressed data.

## Per-Portnum Compression Summary

| Compressor                                    | P1 (TEXT_MESSAGE_APP) | P3 (POSITION_APP) | P4 (NODEINFO_APP) | P5 (ROUTING_APP) | P6 (UNKNOWN) | P65 (STORE_FORWARD_APP) | P66 (UNKNOWN) | P67 (TELEMETRY_APP) | P70 (TRACEROUTE_APP) | P71 (NEIGHBORINFO_APP) | P72 (UNKNOWN) | P278 (UNKNOWN) |
| --------------------------------------------- | --------------------- | ----------------- | ----------------- | ---------------- | ------------ | ----------------------- | ------------- | ------------------- | -------------------- | ---------------------- | ------------- | -------------- |
| `unishox2_alpha_only`                         | 0.6606                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `snowflake_Jorropo`                           | 0.6845                | 1.0000            | 0.7659            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_alpha_num_only`                     | 0.6944                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_no_uni_favor_text`                  | 0.7086                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_favor_alpha`                        | 0.7129                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_no_uni`                             | 0.7138                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_json_no_uni`                        | 0.7141                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_url`                                | 0.7168                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_default`                            | 0.7183                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_xml`                                | 0.7183                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_html`                               | 0.7186                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_json`                               | 0.7186                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_favor_dict`                         | 0.7206                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_alpha_num_sym_only`                 | 0.7222                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_alpha_num_sym_only_text`            | 0.7222                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_favor_sym`                          | 0.7242                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_favor_umlaut`                       | 0.7262                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `unishox2_no_dict`                            | 0.7310                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `meshtasticmodel_V10_EgonElbre`               | 0.7433                | 0.9458            | 0.9686            | 0.7920           | 0.8571       | 0.8553                  | 0.8333        | 0.9496              | 0.9457               | 0.9652                 | 0.9683        | 0.9345         |
| `meshtasticmodel_V8_EgonElbre`                | 0.7433                | 0.9458            | 0.9686            | 0.7920           | 0.8571       | 0.8553                  | 0.8333        | 0.9496              | 0.9457               | 0.9652                 | 0.9683        | 0.9345         |
| `meshtasticmodel_V9_EgonElbre`                | 0.7433                | 0.9458            | 0.9686            | 0.7920           | 0.8571       | 0.8553                  | 0.8333        | 0.9496              | 0.9457               | 0.9652                 | 0.9683        | 0.9345         |
| `meshtasticmodel_V1_EgonElbre`                | 0.7434                | 0.9453            | 0.9660            | 0.7841           | 0.8095       | 0.8553                  | 0.8333        | 0.9502              | 0.9357               | 0.9703                 | 0.9683        | 1.0000         |
| `meshtasticmodel_V4_EgonElbre`                | 0.7434                | 0.9453            | 0.9660            | 0.7841           | 0.8095       | 0.8553                  | 0.8333        | 0.9502              | 0.9357               | 0.9703                 | 0.9683        | 1.0000         |
| `meshtasticmodel_V5_EgonElbre`                | 0.7434                | 0.9453            | 0.9660            | 0.7841           | 0.8095       | 0.8553                  | 0.8333        | 0.9502              | 0.9357               | 0.9703                 | 0.9683        | 1.0000         |
| `meshtasticmodel_V6_EgonElbre`                | 0.7434                | 0.9453            | 0.9660            | 0.7841           | 0.8095       | 0.8553                  | 0.8333        | 0.9502              | 0.9357               | 0.9703                 | 0.9683        | 1.0000         |
| `meshtasticmodel_V7_EgonElbre`                | 0.7436                | 0.9465            | 0.9693            | 0.8129           | 0.8571       | 0.8553                  | 0.8333        | 0.9496              | 0.9471               | 0.9652                 | 0.9683        | 0.9345         |
| `meshtasticmodel_V3_EgonElbre`                | 0.7451                | 0.9498            | 0.9689            | 0.8064           | 0.8571       | 0.8684                  | 0.8333        | 0.9623              | 0.9383               | 0.9779                 | 0.9683        | 1.0000         |
| `meshtasticmodel_V2_EgonElbre`                | 0.8067                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `shoco_TextEn_tmthrgd_Jorropo`                | 0.8089                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `smaz_cespare_Jorropo`                        | 0.8319                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `arithmetic_Tom`                              | 0.8450                | 0.9958            | 0.9887            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 0.9997              | 1.0000               | 0.9414                 | 1.0000        | 0.9709         |
| `shoco_TextEn_tmthrgd`                        | 0.8472                | 1.0000            | 0.9998            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 0.9600         |
| `shoco_WordsEn_tmthrgd_Jorropo`               | 0.8684                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `smaz_cespare`                                | 0.8720                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 0.9855         |
| `shoco_Emails_tmthrgd_Jorropo`                | 0.8901                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `shoco_FilePath_tmthrgd_Jorropo`              | 0.8937                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `shoco_WordsEn_tmthrgd`                       | 0.9047                | 1.0000            | 0.9998            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 0.9600         |
| `shoco_Emails_tmthrgd`                        | 0.9252                | 1.0000            | 0.9998            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 0.9673         |
| `shoco_FilePath_tmthrgd`                      | 0.9281                | 1.0000            | 0.9998            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 0.9636         |
| `flate_klauspost`                             | 0.9324                | 0.9990            | 0.9993            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.7574               | 0.9567                 | 1.0000        | 1.0000         |
| `flate_std`                                   | 0.9410                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.7259               | 0.9711                 | 1.0000        | 1.0000         |
| `meshtasticmodel_pbmodel-o1_EgonElbre`        | 0.9434                | 0.9453            | 0.9617            | 0.7824           | 0.8095       | 0.8553                  | 0.8333        | 0.9496              | 0.9336               | 0.9652                 | 0.9683        | 1.0000         |
| `meshtasticmodel_pbmodel-o2_EgonElbre`        | 0.9434                | 0.9453            | 0.9617            | 0.7824           | 0.8095       | 0.8553                  | 0.8333        | 0.9496              | 0.9336               | 0.9652                 | 0.9683        | 1.0000         |
| `meshtasticmodel_pbmodel-varint-o1_EgonElbre` | 0.9434                | 0.9453            | 0.9617            | 0.7824           | 0.8095       | 0.8553                  | 0.8333        | 0.9496              | 0.9336               | 0.9652                 | 0.9683        | 1.0000         |
| `meshtasticmodel_pbmodel-varint-o2_EgonElbre` | 0.9434                | 0.9453            | 0.9617            | 0.7824           | 0.8095       | 0.8553                  | 0.8333        | 0.9496              | 0.9336               | 0.9652                 | 0.9683        | 1.0000         |
| `meshtasticmodel_pbmodel-varint_EgonElbre`    | 0.9434                | 0.9453            | 0.9617            | 0.7824           | 0.8095       | 0.8553                  | 0.8333        | 0.9496              | 0.9336               | 0.9652                 | 0.9683        | 1.0000         |
| `meshtasticmodel_pbmodel_EgonElbre`           | 0.9434                | 0.9453            | 0.9617            | 0.7824           | 0.8095       | 0.8553                  | 0.8333        | 0.9496              | 0.9336               | 0.9652                 | 0.9683        | 1.0000         |
| `arithmetic`                                  | 0.9450                | 0.9946            | 0.9996            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 0.9999              | 1.0000               | 1.0000                 | 1.0000        | 0.9673         |
| `zlib_klauspost`                              | 0.9470                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.8311               | 0.9788                 | 1.0000        | 1.0000         |
| `lz4_cloudflareHC`                            | 0.9509                | 0.9991            | 0.9999            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.7299               | 0.9720                 | 1.0000        | 1.0000         |
| `zlib_std`                                    | 0.9515                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.7983               | 0.9847                 | 1.0000        | 1.0000         |
| `lz4_cloudflare`                              | 0.9518                | 0.9991            | 0.9999            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.7534               | 0.9771                 | 1.0000        | 1.0000         |
| `lzw_std`                                     | 0.9583                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.7983               | 0.9558                 | 1.0000        | 1.0000         |
| `gzip_klauspost`                              | 0.9601                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.9055               | 0.9992                 | 1.0000        | 1.0000         |
| `gzip_std`                                    | 0.9624                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.8874               | 1.0000                 | 1.0000        | 1.0000         |
| `s2_klauspost`                                | 0.9637                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.8928               | 1.0000                 | 1.0000        | 1.0000         |
| `lz4_pierrec`                                 | 0.9649                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.8881               | 1.0000                 | 1.0000        | 1.0000         |
| `snappy_klauspost`                            | 0.9649                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 0.9182               | 1.0000                 | 1.0000        | 1.0000         |
| `noop`                                        | 1.0000                | 1.0000            | 1.0000            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 1.0000              | 1.0000               | 1.0000                 | 1.0000        | 1.0000         |
| `rle_inkyblackness`                           | 1.0000                | 1.0000            | 0.9999            | 1.0000           | 1.0000       | 1.0000                  | 1.0000        | 0.9999              | 0.8438               | 1.0000                 | 1.0000        | 1.0000         |


## Dictionary/Model Sizes

| Compressor | Dictionary Size (bytes) |
|------------|------------------------|
| `arithmetic` | 7196 |
| `arithmetic_Tom` | 7196 |
| `flate_klauspost` | 0 (algorithm-based) |
| `flate_std` | 0 (algorithm-based) |
| `gzip_klauspost` | 0 (algorithm-based) |
| `gzip_std` | 0 (algorithm-based) |
| `lz4_cloudflare` | 0 (algorithm-based) |
| `lz4_cloudflareHC` | 0 (algorithm-based) |
| `lz4_pierrec` | 0 (algorithm-based) |
| `lzw_std` | 0 (algorithm-based) |
| `meshtasticmodel_V10_EgonElbre` | 4096 |
| `meshtasticmodel_V1_EgonElbre` | 4096 |
| `meshtasticmodel_V2_EgonElbre` | 4096 |
| `meshtasticmodel_V3_EgonElbre` | 4096 |
| `meshtasticmodel_V4_EgonElbre` | 4096 |
| `meshtasticmodel_V5_EgonElbre` | 4096 |
| `meshtasticmodel_V6_EgonElbre` | 4096 |
| `meshtasticmodel_V7_EgonElbre` | 4096 |
| `meshtasticmodel_V8_EgonElbre` | 4096 |
| `meshtasticmodel_V9_EgonElbre` | 4096 |
| `meshtasticmodel_pbmodel-o1_EgonElbre` | 8192 |
| `meshtasticmodel_pbmodel-o2_EgonElbre` | 8192 |
| `meshtasticmodel_pbmodel-varint-o1_EgonElbre` | 8192 |
| `meshtasticmodel_pbmodel-varint-o2_EgonElbre` | 8192 |
| `meshtasticmodel_pbmodel-varint_EgonElbre` | 8192 |
| `meshtasticmodel_pbmodel_EgonElbre` | 8192 |
| `noop` | 0 (algorithm-based) |
| `rle_inkyblackness` | 0 (algorithm-based) |
| `s2_klauspost` | 0 (algorithm-based) |
| `shoco_Emails_tmthrgd` | 1024 |
| `shoco_Emails_tmthrgd_Jorropo` | 1024 |
| `shoco_FilePath_tmthrgd` | 2048 |
| `shoco_FilePath_tmthrgd_Jorropo` | 2048 |
| `shoco_TextEn_tmthrgd` | 4096 |
| `shoco_TextEn_tmthrgd_Jorropo` | 4096 |
| `shoco_WordsEn_tmthrgd` | 3072 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 3072 |
| `smaz_cespare` | 1280 |
| `smaz_cespare_Jorropo` | 1280 |
| `snappy_klauspost` | 0 (algorithm-based) |
| `snowflake_Jorropo` | 512 |
| `unishox2_alpha_num_only` | 0 (algorithm-based) |
| `unishox2_alpha_num_sym_only` | 0 (algorithm-based) |
| `unishox2_alpha_num_sym_only_text` | 0 (algorithm-based) |
| `unishox2_alpha_only` | 0 (algorithm-based) |
| `unishox2_default` | 0 (algorithm-based) |
| `unishox2_favor_alpha` | 0 (algorithm-based) |
| `unishox2_favor_dict` | 0 (algorithm-based) |
| `unishox2_favor_sym` | 0 (algorithm-based) |
| `unishox2_favor_umlaut` | 0 (algorithm-based) |
| `unishox2_html` | 0 (algorithm-based) |
| `unishox2_json` | 0 (algorithm-based) |
| `unishox2_json_no_uni` | 0 (algorithm-based) |
| `unishox2_no_dict` | 0 (algorithm-based) |
| `unishox2_no_uni` | 0 (algorithm-based) |
| `unishox2_no_uni_favor_text` | 0 (algorithm-based) |
| `unishox2_url` | 0 (algorithm-based) |
| `unishox2_xml` | 0 (algorithm-based) |
| `zlib_klauspost` | 0 (algorithm-based) |
| `zlib_std` | 0 (algorithm-based) |

## Results

| Compressor | Average Reciprocal Compression Ratio (TEXT_MESSAGE_APP only) |
|------------|--------------------------------------------------------------|
| `unishox2_alpha_only` | 0.6606 |
| `snowflake_Jorropo` | 0.6845 |
| `unishox2_alpha_num_only` | 0.6944 |
| `unishox2_no_uni_favor_text` | 0.7086 |
| `unishox2_favor_alpha` | 0.7129 |
| `unishox2_no_uni` | 0.7138 |
| `unishox2_json_no_uni` | 0.7141 |
| `unishox2_url` | 0.7168 |
| `unishox2_default` | 0.7183 |
| `unishox2_xml` | 0.7183 |
| `unishox2_html` | 0.7186 |
| `unishox2_json` | 0.7186 |
| `unishox2_favor_dict` | 0.7206 |
| `unishox2_alpha_num_sym_only` | 0.7222 |
| `unishox2_alpha_num_sym_only_text` | 0.7222 |
| `unishox2_favor_sym` | 0.7242 |
| `unishox2_favor_umlaut` | 0.7262 |
| `unishox2_no_dict` | 0.7310 |
| `meshtasticmodel_V10_EgonElbre` | 0.7433 |
| `meshtasticmodel_V8_EgonElbre` | 0.7433 |
| `meshtasticmodel_V9_EgonElbre` | 0.7433 |
| `meshtasticmodel_V1_EgonElbre` | 0.7434 |
| `meshtasticmodel_V4_EgonElbre` | 0.7434 |
| `meshtasticmodel_V5_EgonElbre` | 0.7434 |
| `meshtasticmodel_V6_EgonElbre` | 0.7434 |
| `meshtasticmodel_V7_EgonElbre` | 0.7436 |
| `meshtasticmodel_V3_EgonElbre` | 0.7451 |
| `meshtasticmodel_V2_EgonElbre` | 0.8067 |
| `shoco_TextEn_tmthrgd_Jorropo` | 0.8089 |
| `smaz_cespare_Jorropo` | 0.8319 |
| `arithmetic_Tom` | 0.8450 |
| `shoco_TextEn_tmthrgd` | 0.8472 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.8684 |
| `smaz_cespare` | 0.8720 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.8901 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.8937 |
| `shoco_WordsEn_tmthrgd` | 0.9047 |
| `shoco_Emails_tmthrgd` | 0.9252 |
| `shoco_FilePath_tmthrgd` | 0.9281 |
| `flate_klauspost` | 0.9324 |
| `flate_std` | 0.9410 |
| `meshtasticmodel_pbmodel-o1_EgonElbre` | 0.9434 |
| `meshtasticmodel_pbmodel-o2_EgonElbre` | 0.9434 |
| `meshtasticmodel_pbmodel-varint-o1_EgonElbre` | 0.9434 |
| `meshtasticmodel_pbmodel-varint-o2_EgonElbre` | 0.9434 |
| `meshtasticmodel_pbmodel-varint_EgonElbre` | 0.9434 |
| `meshtasticmodel_pbmodel_EgonElbre` | 0.9434 |
| `arithmetic` | 0.9450 |
| `zlib_klauspost` | 0.9470 |
| `lz4_cloudflareHC` | 0.9509 |
| `zlib_std` | 0.9515 |
| `lz4_cloudflare` | 0.9518 |
| `lzw_std` | 0.9583 |
| `gzip_klauspost` | 0.9601 |
| `gzip_std` | 0.9624 |
| `s2_klauspost` | 0.9637 |
| `lz4_pierrec` | 0.9649 |
| `snappy_klauspost` | 0.9649 |
| `noop` | 1.0000 |
| `rle_inkyblackness` | 1.0000 |

| Compressor | Average Reciprocal Compression Ratio |
|------------|--------------------------------------|
| `snowflake_Jorropo` | 0.8790 |
| `meshtasticmodel_V1_EgonElbre` | 0.9524 |
| `meshtasticmodel_V4_EgonElbre` | 0.9524 |
| `meshtasticmodel_V5_EgonElbre` | 0.9524 |
| `meshtasticmodel_V6_EgonElbre` | 0.9524 |
| `meshtasticmodel_pbmodel-o1_EgonElbre` | 0.9525 |
| `meshtasticmodel_pbmodel-o2_EgonElbre` | 0.9525 |
| `meshtasticmodel_pbmodel-varint-o1_EgonElbre` | 0.9525 |
| `meshtasticmodel_pbmodel-varint-o2_EgonElbre` | 0.9525 |
| `meshtasticmodel_pbmodel-varint_EgonElbre` | 0.9525 |
| `meshtasticmodel_pbmodel_EgonElbre` | 0.9525 |
| `meshtasticmodel_V10_EgonElbre` | 0.9537 |
| `meshtasticmodel_V8_EgonElbre` | 0.9537 |
| `meshtasticmodel_V9_EgonElbre` | 0.9537 |
| `meshtasticmodel_V7_EgonElbre` | 0.9545 |
| `meshtasticmodel_V3_EgonElbre` | 0.9580 |
| `arithmetic_Tom` | 0.9912 |
| `unishox2_alpha_only` | 0.9958 |
| `unishox2_alpha_num_only` | 0.9962 |
| `unishox2_no_uni_favor_text` | 0.9964 |
| `unishox2_favor_alpha` | 0.9965 |
| `unishox2_no_uni` | 0.9965 |
| `unishox2_json_no_uni` | 0.9965 |
| `unishox2_url` | 0.9965 |
| `unishox2_default` | 0.9965 |
| `unishox2_xml` | 0.9965 |
| `unishox2_html` | 0.9965 |
| `unishox2_json` | 0.9965 |
| `unishox2_favor_dict` | 0.9966 |
| `unishox2_alpha_num_sym_only` | 0.9966 |
| `unishox2_alpha_num_sym_only_text` | 0.9966 |
| `unishox2_favor_sym` | 0.9966 |
| `unishox2_favor_umlaut` | 0.9966 |
| `unishox2_no_dict` | 0.9967 |
| `meshtasticmodel_V2_EgonElbre` | 0.9976 |
| `shoco_TextEn_tmthrgd_Jorropo` | 0.9976 |
| `flate_klauspost` | 0.9976 |
| `arithmetic` | 0.9977 |
| `smaz_cespare_Jorropo` | 0.9979 |
| `shoco_TextEn_tmthrgd` | 0.9980 |
| `lz4_cloudflareHC` | 0.9982 |
| `flate_std` | 0.9983 |
| `lz4_cloudflare` | 0.9983 |
| `shoco_WordsEn_tmthrgd_Jorropo` | 0.9984 |
| `smaz_cespare` | 0.9984 |
| `shoco_Emails_tmthrgd_Jorropo` | 0.9986 |
| `shoco_FilePath_tmthrgd_Jorropo` | 0.9987 |
| `lzw_std` | 0.9987 |
| `zlib_std` | 0.9987 |
| `shoco_WordsEn_tmthrgd` | 0.9987 |
| `zlib_klauspost` | 0.9987 |
| `shoco_Emails_tmthrgd` | 0.9990 |
| `shoco_FilePath_tmthrgd` | 0.9990 |
| `gzip_std` | 0.9992 |
| `gzip_klauspost` | 0.9992 |
| `lz4_pierrec` | 0.9992 |
| `s2_klauspost` | 0.9992 |
| `snappy_klauspost` | 0.9993 |
| `rle_inkyblackness` | 0.9994 |
| `noop` | 1.0000 |

## CDF Graphs

The following graphs show the cumulative distribution function (CDF) of the reciprocal compression ratios for each compressor.

### `snowflake_Jorropo`

![snowflake_Jorropo only TEXT_MESSAGE_APP CDF](graphs/snowflake_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![snowflake_Jorropo CDF](graphs/snowflake_Jorropo_cdf.png)

### `meshtasticmodel_V1_EgonElbre`

![meshtasticmodel_V1_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V1_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V1_EgonElbre CDF](graphs/meshtasticmodel_V1_EgonElbre_cdf.png)

### `meshtasticmodel_V4_EgonElbre`

![meshtasticmodel_V4_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V4_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V4_EgonElbre CDF](graphs/meshtasticmodel_V4_EgonElbre_cdf.png)

### `meshtasticmodel_V5_EgonElbre`

![meshtasticmodel_V5_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V5_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V5_EgonElbre CDF](graphs/meshtasticmodel_V5_EgonElbre_cdf.png)

### `meshtasticmodel_V6_EgonElbre`

![meshtasticmodel_V6_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V6_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V6_EgonElbre CDF](graphs/meshtasticmodel_V6_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-o1_EgonElbre`

![meshtasticmodel_pbmodel-o1_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-o1_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-o1_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-o1_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-o2_EgonElbre`

![meshtasticmodel_pbmodel-o2_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-o2_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-o2_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-o2_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-varint-o1_EgonElbre`

![meshtasticmodel_pbmodel-varint-o1_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-varint-o1_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-varint-o1_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-varint-o1_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-varint-o2_EgonElbre`

![meshtasticmodel_pbmodel-varint-o2_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-varint-o2_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-varint-o2_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-varint-o2_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel-varint_EgonElbre`

![meshtasticmodel_pbmodel-varint_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel-varint_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel-varint_EgonElbre CDF](graphs/meshtasticmodel_pbmodel-varint_EgonElbre_cdf.png)

### `meshtasticmodel_pbmodel_EgonElbre`

![meshtasticmodel_pbmodel_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_pbmodel_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_pbmodel_EgonElbre CDF](graphs/meshtasticmodel_pbmodel_EgonElbre_cdf.png)

### `meshtasticmodel_V10_EgonElbre`

![meshtasticmodel_V10_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V10_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V10_EgonElbre CDF](graphs/meshtasticmodel_V10_EgonElbre_cdf.png)

### `meshtasticmodel_V8_EgonElbre`

![meshtasticmodel_V8_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V8_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V8_EgonElbre CDF](graphs/meshtasticmodel_V8_EgonElbre_cdf.png)

### `meshtasticmodel_V9_EgonElbre`

![meshtasticmodel_V9_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V9_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V9_EgonElbre CDF](graphs/meshtasticmodel_V9_EgonElbre_cdf.png)

### `meshtasticmodel_V7_EgonElbre`

![meshtasticmodel_V7_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V7_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V7_EgonElbre CDF](graphs/meshtasticmodel_V7_EgonElbre_cdf.png)

### `meshtasticmodel_V3_EgonElbre`

![meshtasticmodel_V3_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V3_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V3_EgonElbre CDF](graphs/meshtasticmodel_V3_EgonElbre_cdf.png)

### `arithmetic_Tom`

![arithmetic_Tom only TEXT_MESSAGE_APP CDF](graphs/arithmetic_Tom_only_TEXT_MESSAGE_APP_cdf.png)

![arithmetic_Tom CDF](graphs/arithmetic_Tom_cdf.png)

### `unishox2_alpha_only`

![unishox2_alpha_only only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_only_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_only CDF](graphs/unishox2_alpha_only_cdf.png)

### `unishox2_alpha_num_only`

![unishox2_alpha_num_only only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_num_only_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_num_only CDF](graphs/unishox2_alpha_num_only_cdf.png)

### `unishox2_no_uni_favor_text`

![unishox2_no_uni_favor_text only TEXT_MESSAGE_APP CDF](graphs/unishox2_no_uni_favor_text_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_no_uni_favor_text CDF](graphs/unishox2_no_uni_favor_text_cdf.png)

### `unishox2_favor_alpha`

![unishox2_favor_alpha only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_alpha_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_alpha CDF](graphs/unishox2_favor_alpha_cdf.png)

### `unishox2_no_uni`

![unishox2_no_uni only TEXT_MESSAGE_APP CDF](graphs/unishox2_no_uni_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_no_uni CDF](graphs/unishox2_no_uni_cdf.png)

### `unishox2_json_no_uni`

![unishox2_json_no_uni only TEXT_MESSAGE_APP CDF](graphs/unishox2_json_no_uni_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_json_no_uni CDF](graphs/unishox2_json_no_uni_cdf.png)

### `unishox2_url`

![unishox2_url only TEXT_MESSAGE_APP CDF](graphs/unishox2_url_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_url CDF](graphs/unishox2_url_cdf.png)

### `unishox2_default`

![unishox2_default only TEXT_MESSAGE_APP CDF](graphs/unishox2_default_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_default CDF](graphs/unishox2_default_cdf.png)

### `unishox2_xml`

![unishox2_xml only TEXT_MESSAGE_APP CDF](graphs/unishox2_xml_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_xml CDF](graphs/unishox2_xml_cdf.png)

### `unishox2_html`

![unishox2_html only TEXT_MESSAGE_APP CDF](graphs/unishox2_html_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_html CDF](graphs/unishox2_html_cdf.png)

### `unishox2_json`

![unishox2_json only TEXT_MESSAGE_APP CDF](graphs/unishox2_json_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_json CDF](graphs/unishox2_json_cdf.png)

### `unishox2_favor_dict`

![unishox2_favor_dict only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_dict_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_dict CDF](graphs/unishox2_favor_dict_cdf.png)

### `unishox2_alpha_num_sym_only`

![unishox2_alpha_num_sym_only only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_num_sym_only_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_num_sym_only CDF](graphs/unishox2_alpha_num_sym_only_cdf.png)

### `unishox2_alpha_num_sym_only_text`

![unishox2_alpha_num_sym_only_text only TEXT_MESSAGE_APP CDF](graphs/unishox2_alpha_num_sym_only_text_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_alpha_num_sym_only_text CDF](graphs/unishox2_alpha_num_sym_only_text_cdf.png)

### `unishox2_favor_sym`

![unishox2_favor_sym only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_sym_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_sym CDF](graphs/unishox2_favor_sym_cdf.png)

### `unishox2_favor_umlaut`

![unishox2_favor_umlaut only TEXT_MESSAGE_APP CDF](graphs/unishox2_favor_umlaut_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_favor_umlaut CDF](graphs/unishox2_favor_umlaut_cdf.png)

### `unishox2_no_dict`

![unishox2_no_dict only TEXT_MESSAGE_APP CDF](graphs/unishox2_no_dict_only_TEXT_MESSAGE_APP_cdf.png)

![unishox2_no_dict CDF](graphs/unishox2_no_dict_cdf.png)

### `meshtasticmodel_V2_EgonElbre`

![meshtasticmodel_V2_EgonElbre only TEXT_MESSAGE_APP CDF](graphs/meshtasticmodel_V2_EgonElbre_only_TEXT_MESSAGE_APP_cdf.png)

![meshtasticmodel_V2_EgonElbre CDF](graphs/meshtasticmodel_V2_EgonElbre_cdf.png)

### `shoco_TextEn_tmthrgd_Jorropo`

![shoco_TextEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_TextEn_tmthrgd_Jorropo CDF](graphs/shoco_TextEn_tmthrgd_Jorropo_cdf.png)

### `flate_klauspost`

![flate_klauspost only TEXT_MESSAGE_APP CDF](graphs/flate_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![flate_klauspost CDF](graphs/flate_klauspost_cdf.png)

### `arithmetic`

![arithmetic only TEXT_MESSAGE_APP CDF](graphs/arithmetic_only_TEXT_MESSAGE_APP_cdf.png)

![arithmetic CDF](graphs/arithmetic_cdf.png)

### `smaz_cespare_Jorropo`

![smaz_cespare_Jorropo only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![smaz_cespare_Jorropo CDF](graphs/smaz_cespare_Jorropo_cdf.png)

### `shoco_TextEn_tmthrgd`

![shoco_TextEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_TextEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_TextEn_tmthrgd CDF](graphs/shoco_TextEn_tmthrgd_cdf.png)

### `lz4_cloudflareHC`

![lz4_cloudflareHC only TEXT_MESSAGE_APP CDF](graphs/lz4_cloudflareHC_only_TEXT_MESSAGE_APP_cdf.png)

![lz4_cloudflareHC CDF](graphs/lz4_cloudflareHC_cdf.png)

### `flate_std`

![flate_std only TEXT_MESSAGE_APP CDF](graphs/flate_std_only_TEXT_MESSAGE_APP_cdf.png)

![flate_std CDF](graphs/flate_std_cdf.png)

### `lz4_cloudflare`

![lz4_cloudflare only TEXT_MESSAGE_APP CDF](graphs/lz4_cloudflare_only_TEXT_MESSAGE_APP_cdf.png)

![lz4_cloudflare CDF](graphs/lz4_cloudflare_cdf.png)

### `shoco_WordsEn_tmthrgd_Jorropo`

![shoco_WordsEn_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_WordsEn_tmthrgd_Jorropo CDF](graphs/shoco_WordsEn_tmthrgd_Jorropo_cdf.png)

### `smaz_cespare`

![smaz_cespare only TEXT_MESSAGE_APP CDF](graphs/smaz_cespare_only_TEXT_MESSAGE_APP_cdf.png)

![smaz_cespare CDF](graphs/smaz_cespare_cdf.png)

### `shoco_Emails_tmthrgd_Jorropo`

![shoco_Emails_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_Emails_tmthrgd_Jorropo CDF](graphs/shoco_Emails_tmthrgd_Jorropo_cdf.png)

### `shoco_FilePath_tmthrgd_Jorropo`

![shoco_FilePath_tmthrgd_Jorropo only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_FilePath_tmthrgd_Jorropo CDF](graphs/shoco_FilePath_tmthrgd_Jorropo_cdf.png)

### `lzw_std`

![lzw_std only TEXT_MESSAGE_APP CDF](graphs/lzw_std_only_TEXT_MESSAGE_APP_cdf.png)

![lzw_std CDF](graphs/lzw_std_cdf.png)

### `zlib_std`

![zlib_std only TEXT_MESSAGE_APP CDF](graphs/zlib_std_only_TEXT_MESSAGE_APP_cdf.png)

![zlib_std CDF](graphs/zlib_std_cdf.png)

### `shoco_WordsEn_tmthrgd`

![shoco_WordsEn_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_WordsEn_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_WordsEn_tmthrgd CDF](graphs/shoco_WordsEn_tmthrgd_cdf.png)

### `zlib_klauspost`

![zlib_klauspost only TEXT_MESSAGE_APP CDF](graphs/zlib_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![zlib_klauspost CDF](graphs/zlib_klauspost_cdf.png)

### `shoco_Emails_tmthrgd`

![shoco_Emails_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_Emails_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_Emails_tmthrgd CDF](graphs/shoco_Emails_tmthrgd_cdf.png)

### `shoco_FilePath_tmthrgd`

![shoco_FilePath_tmthrgd only TEXT_MESSAGE_APP CDF](graphs/shoco_FilePath_tmthrgd_only_TEXT_MESSAGE_APP_cdf.png)

![shoco_FilePath_tmthrgd CDF](graphs/shoco_FilePath_tmthrgd_cdf.png)

### `gzip_std`

![gzip_std only TEXT_MESSAGE_APP CDF](graphs/gzip_std_only_TEXT_MESSAGE_APP_cdf.png)

![gzip_std CDF](graphs/gzip_std_cdf.png)

### `gzip_klauspost`

![gzip_klauspost only TEXT_MESSAGE_APP CDF](graphs/gzip_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![gzip_klauspost CDF](graphs/gzip_klauspost_cdf.png)

### `lz4_pierrec`

![lz4_pierrec only TEXT_MESSAGE_APP CDF](graphs/lz4_pierrec_only_TEXT_MESSAGE_APP_cdf.png)

![lz4_pierrec CDF](graphs/lz4_pierrec_cdf.png)

### `s2_klauspost`

![s2_klauspost only TEXT_MESSAGE_APP CDF](graphs/s2_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![s2_klauspost CDF](graphs/s2_klauspost_cdf.png)

### `snappy_klauspost`

![snappy_klauspost only TEXT_MESSAGE_APP CDF](graphs/snappy_klauspost_only_TEXT_MESSAGE_APP_cdf.png)

![snappy_klauspost CDF](graphs/snappy_klauspost_cdf.png)

### `rle_inkyblackness`

![rle_inkyblackness only TEXT_MESSAGE_APP CDF](graphs/rle_inkyblackness_only_TEXT_MESSAGE_APP_cdf.png)

![rle_inkyblackness CDF](graphs/rle_inkyblackness_cdf.png)

### `noop`

![noop only TEXT_MESSAGE_APP CDF](graphs/noop_only_TEXT_MESSAGE_APP_cdf.png)

![noop CDF](graphs/noop_cdf.png)

