#!/bin/bash

APP_VER=$GO_PIPELINE_LABEL
DEPLOY_STACK_NAME="$ECS_STACK_NAME-deploy"
AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
AWS_DEFAULT_REGION=$AWS_DEFAULT_REGION
AWS_ACCOUNT_ID=$AWS_ACCOUNT_ID

echo "DEPLOY_STACK_NAME -> $DEPLOY_STACK_NAME"
echo "ECS_STACK_NAME -> $ECS_STACK_NAME"
echo "BASE_STACK_NAME -> $BASE_STACK_NAME"
echo "BUCKET_NAME -> $BUCKET_NAME"
echo "SERVICE_NAME -> $SERVICE_NAME"
echo "APP_VER -> $APP_VER"
echo "ENV -> $ENV"

echo "Deploying application into ECS ..."
aws cloudformation describe-stacks --stack-name $DEPLOY_STACK_NAME
isExist=$?

if [ $isExist -ne 0 ]
then

  echo "Createing new stack -> $DEPLOY_STACK_NAME"
  aws cloudformation create-stack --stack-name $DEPLOY_STACK_NAME \
    --template-url `aws s3 presign s3://$BUCKET_NAME/ecs/app-main.yaml`  \
    --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM \
    --parameters \
    ParameterKey=baseStackName,ParameterValue=$BASE_STACK_NAME \
    ParameterKey=ecsStackName,ParameterValue=$ECS_STACK_NAME \
    ParameterKey=env,ParameterValue=$ENV \
    ParameterKey=imageVersion,ParameterValue=$APP_VER \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME
  isExist=$?

  if [ $isExist -eq 0 ]
  then
    aws cloudformation wait stack-create-complete --stack-name $DEPLOY_STACK_NAME
    isExist=$?
  fi

else

  echo "Updating new stack -> $DEPLOY_STACK_NAME"
  aws cloudformation update-stack --stack-name $DEPLOY_STACK_NAME \
    --template-url `aws s3 presign s3://$BUCKET_NAME/ecs/app-main.yaml`  \
    --parameters \
    ParameterKey=baseStackName,ParameterValue=$BASE_STACK_NAME \
    ParameterKey=ecsStackName,ParameterValue=$ECS_STACK_NAME \
    ParameterKey=env,ParameterValue=$ENV \
    ParameterKey=imageVersion,ParameterValue=$APP_VER \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME
  isExist=$?

  if [ $isExist -eq 0 ]
  then
    aws cloudformation wait stack-update-complete --stack-name $DEPLOY_STACK_NAME
    isExist=$?
  fi

fi

if [ $isExist -eq 0 ]
then
  echo `aws cloudformation describe-stacks --stack-name $DEPLOY_STACK_NAME` > output.json
  cat output.json |jq
  echo "Click URL -> http://"`cat output.json |jq '.Stacks[0].Outputs[].OutputValue' |grep -v "arn:aws"|sed -e "s/\"//g"`":8080/encode" > to.mail
  aws sns publish --topic-arn "arn:aws:sns:us-east-1:530820415924:cicd-notification" \
    --subject  "Master, Check Out the Result of Deplouyment. `date`" \
    --message file://to.mail

  echo "Done"
else
  echo "Operation was failed."
  echo "Deployment was falied and contact with administrator, please!!!" > to.mail
  aws sns publish --topic-arn "arn:aws:sns:us-east-1:530820415924:cicd-notification" \
    --subject  "Master, deplouyment was falied. `date`" \
    --message file://to.mail
  exit 255
fi
