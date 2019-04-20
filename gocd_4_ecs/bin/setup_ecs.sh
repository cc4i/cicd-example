#!/bin/bash


echo "ECS_STACK_NAME -> $ECS_STACK_NAME"
echo "BUCKET_NAME -> $BUCKET_NAME"
echo "SERVICE_NAME -> $SERVICE_NAME"

echo "Seting ECS  environment ..."
aws cloudformation describe-stacks --stack-name $ECS_STACK_NAME
isExist=$?

if [ $isExist -ne 0 ]
then

  echo "Createing new stack -> $ECS_STACK_NAME"
  aws cloudformation create-stack --stack-name $ECS_STACK_NAME \
    --template-url `aws s3 presign s3://$BUCKET_NAME/ecs/app-ecs.yaml`  \
    --parameters \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME
  isExist=$?

  if [ $isExist -eq 0 ]
  then
    aws cloudformation wait stack-create-complete --stack-name $ECS_STACK_NAME
  fi

else

  echo "Updating new stack -> $ECS_STACK_NAME"
  aws cloudformation update-stack --stack-name $ECS_STACK_NAME \
    --template-url `aws s3 presign s3://$BUCKET_NAME/ecs/app-ecs.yaml`  \
    --parameters \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME
  isExist=$?

  if [ $isExist -eq 0 ]
  then
    aws cloudformation wait stack-update-complete --stack-name $ECS_STACK_NAME
  fi

fi
echo "Done"
