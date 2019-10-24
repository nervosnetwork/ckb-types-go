// Reference: https://github.com/yangby-cryptape/ckb-ffi

#ifndef CKB_FFI_H
#define CKB_FFI_H

#ifdef __cplusplus
#define _CPP_BEGIN extern "C" {
#define _CPP_END }
_CPP_BEGIN
#endif /* __cplusplus */

#include <stdint.h>

#define CKB_FFI_OK                              0x00
#define CKB_FFI_INVALID_JSON                    0x01
#define CKB_FFI_INVALID_MOLECULE                0x02
#define CKB_FFI_UNSUPPORTED_TYPE                0x03

typedef struct {
    uint64_t len;
    uint8_t *data;
} buffer_t;

extern void buffer_free(buffer_t);

extern int32_t ckb_encode(buffer_t* output_tx, char* type_name, char* json_tx);
extern int32_t ckb_decode(buffer_t* output_tx, char* type_name, buffer_t* mol_tx);

#ifdef __cplusplus
_CPP_END
#undef _CPP_BEGIN
#undef _CPP_END
#endif /* __cplusplus */

#endif /* CKB_FFI_H */
