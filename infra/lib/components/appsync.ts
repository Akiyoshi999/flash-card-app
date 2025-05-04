import { Construct } from "constructs";
import * as appsync from "aws-cdk-lib/aws-appsync";
import { EnvProps } from "../type";
import * as path from "path";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";
import * as cognito from "aws-cdk-lib/aws-cognito";
interface AppsyncConstructProps extends EnvProps {
  table: dynamodb.Table;
  userPool: cognito.UserPool;
}

export class AppsyncConstruct extends Construct {
  public readonly api: appsync.GraphqlApi;
  constructor(scope: Construct, id: string, props: AppsyncConstructProps) {
    super(scope, id);

    this.api = new appsync.GraphqlApi(
      this,
      `Appsync-${props.stackNameSuffix}`,
      {
        name: `FlashCardApp-${props.stackNameSuffix}-API`,
        definition: appsync.Definition.fromSchema(
          appsync.SchemaFile.fromAsset(
            path.join(__dirname, "../../../backend/appsync/schema.graphql")
          )
        ),
        authorizationConfig: {
          defaultAuthorization: {
            authorizationType: appsync.AuthorizationType.USER_POOL,
            userPoolConfig: {
              userPool: props.userPool,
            },
          },
          additionalAuthorizationModes: [
            {
              authorizationType: appsync.AuthorizationType.API_KEY,
            },
          ],
        },
      }
    );

    const dataSource = this.api.addDynamoDbDataSource(
      `DynamoDB-${props.stackNameSuffix}`,
      props.table
    );

    this.api.createResolver("createDeck", {
      typeName: "Mutation",
      fieldName: "createDeck",
      dataSource: dataSource,
      requestMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(
          __dirname,
          "../../../backend/appsync/Deck/Mutation/createDeck.req.vtl"
        )
      ),
      responseMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(
          __dirname,
          "../../../backend/appsync/Deck/Mutation/createDeck.res.vtl"
        )
      ),
    });
  }
}
