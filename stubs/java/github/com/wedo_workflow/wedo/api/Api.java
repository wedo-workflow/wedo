// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: api/api.proto

package github.com.wedo_workflow.wedo.api;

public final class Api {
  private Api() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\rapi/api.proto\022!github.com.wedo_workflo" +
      "w.wedo.api\032\034google/api/annotations.proto" +
      "\032\037google/protobuf/timestamp.proto\032\030api/a" +
      "pi_deployment.proto2\305\001\n\026CloudTaskManageS" +
      "ervice\022\252\001\n\020DeploymentCreate\022:.github.com" +
      ".wedo_workflow.wedo.api.DeploymentCreate" +
      "Request\032;.github.com.wedo_workflow.wedo." +
      "api.DeploymentCreateResponse\"\035\202\323\344\223\002\027\"\022/d" +
      "eployment/create:\001*B#Z!github.com/wedo-w" +
      "orkflow/wedo/apib\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.api.AnnotationsProto.getDescriptor(),
          com.google.protobuf.TimestampProto.getDescriptor(),
          github.com.wedo_workflow.wedo.api.ApiDeployment.getDescriptor(),
        });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.api.AnnotationsProto.http);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.api.AnnotationsProto.getDescriptor();
    com.google.protobuf.TimestampProto.getDescriptor();
    github.com.wedo_workflow.wedo.api.ApiDeployment.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
