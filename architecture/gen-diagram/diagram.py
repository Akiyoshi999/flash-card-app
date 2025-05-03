import os
from diagrams import Diagram
from diagrams.aws.compute import Lambda
from diagrams.aws.database import Dynamodb
from diagrams.aws.security import Cognito
from diagrams.aws.mobile import Appsync
from diagrams.onprem.client import User


output_dir = os.path.join(os.path.dirname(__file__), "../diagrams/flash-card-app")


with Diagram("Flash Card App Architecture", show=False, filename=output_dir):
    user = User("User")
    cognito = Cognito("Cognito UserPool")
    appsync = Appsync("AppSync API")
    ddb = Dynamodb("DynamoDB Table")
    lambda_bff = Lambda("Lambda BFF")

    user >> cognito >> appsync
    appsync >> lambda_bff >> ddb
