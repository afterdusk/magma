add_compile_options(-std=c++14)

set(N11_LTE_CPP_PROTOS apn policydb mobilityd subscriberdb session_manager subscriberauth)
set(N11_LTE_GRPC_PROTOS session_manager mobilityd subscriberauth)
set(N11_ORC8R_GRPC_PROTOS common)

list(APPEND PROTO_SRCS "")
list(APPEND PROTO_HDRS "")

generate_all_protos("${N11_LTE_CPP_PROTOS}" "${N11_ORC8R_GRPC_PROTOS}"
    "${N11_LTE_GRPC_PROTOS}" "" "${PROTO_SRCS}" "${PROTO_HDRS}")



add_library(LIB_N11
    SmfServiceClient.cpp
    amf_client_proto_msg_to_itti_msg.cpp
    M5GAuthenticationServiceClient.cpp
    M5GMobilityServiceClient.cpp
    ${PROTO_SRCS}
    ${PROTO_HDRS}
    )


target_link_libraries(LIB_N11
    COMMON
    LIB_MOBILITY_CLIENT
    SERVICE_REGISTRY
    ASYNC_GRPC
    MAGMA_CONFIG
    )

target_include_directories(LIB_N11 PUBLIC
    ${CMAKE_CURRENT_SOURCE_DIR}
    $<TARGET_FILE_DIR:LIB_N11>
    )
