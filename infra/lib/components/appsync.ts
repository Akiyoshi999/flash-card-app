import { Construct } from "constructs";
import * as appsync from "aws-cdk-lib/aws-appsync";
import { EnvProps } from "../type";
import * as path from "path";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";
import * as cognito from "aws-cdk-lib/aws-cognito";
import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";

interface AppsyncConstructProps extends EnvProps {
  table: dynamodb.Table;
  userPool: cognito.UserPool;
  bffFunction: GoFunction;
}

export class AppsyncConstruct extends Construct {
  public readonly api: appsync.GraphqlApi;
  constructor(scope: Construct, id: string, props: AppsyncConstructProps) {
    super(scope, id);

    const { bffFunction } = props;

    // スキーマを定義
    this.api = new appsync.GraphqlApi(
      this,
      `Appsync-${props.stackNameSuffix}`,
      {
        name: `FlashCardApp-${props.stackNameSuffix}-API`,
        definition: appsync.Definition.fromSchema(
          appsync.SchemaFile.fromAsset(
            path.join(__dirname, "../../schema/schema.graphql")
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

    const lambdaDataSource = this.api.addLambdaDataSource(
      `Lambda-${props.stackNameSuffix}`,
      bffFunction
    );

    // createDeckリゾルバを作成
    this.api.createResolver("createDeck", {
      typeName: "Mutation",
      fieldName: "createDeck",
      dataSource: lambdaDataSource,
      requestMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(
          __dirname,
          "../../appsync-vtl/Deck/Mutation/createDeck/req.vtl"
        )
      ),
      responseMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(
          __dirname,
          "../../appsync-vtl/Deck/Mutation/createDeck/res.vtl"
        )
      ),
    });

    // deleteDeckリゾルバを作成
    this.api.createResolver("deleteDeck", {
      typeName: "Mutation",
      fieldName: "deleteDeck",
      dataSource: lambdaDataSource,
      requestMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(
          __dirname,
          "../../appsync-vtl/Deck/Mutation/deleteDeck/req.vtl"
        )
      ),
      responseMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(
          __dirname,
          "../../appsync-vtl/Deck/Mutation/deleteDeck/res.vtl"
        )
      ),
    });

    // getDeckリゾルバを作成
    this.api.createResolver("getDeck", {
      typeName: "Query",
      fieldName: "getDeck",
      dataSource: lambdaDataSource,
      requestMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(__dirname, "../../appsync-vtl/Deck/Query/getDeck/req.vtl")
      ),
      responseMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(__dirname, "../../appsync-vtl/Deck/Query/getDeck/res.vtl")
      ),
    });

    // listDecksリゾルバを作成
    this.api.createResolver("listDecks", {
      typeName: "Query",
      fieldName: "listDecks",
      dataSource: lambdaDataSource,
      requestMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(__dirname, "../../appsync-vtl/Deck/Query/listDecks/req.vtl")
      ),
      responseMappingTemplate: appsync.MappingTemplate.fromFile(
        path.join(__dirname, "../../appsync-vtl/Deck/Query/listDecks/res.vtl")
      ),
    });
  }
}
