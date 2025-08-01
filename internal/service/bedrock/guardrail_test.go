// // Copyright (c) HashiCorp, Inc.
// // SPDX-License-Identifier: MPL-2.0

package bedrock_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/endpoints"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	tfbedrock "github.com/hashicorp/terraform-provider-aws/internal/service/bedrock"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccBedrockGuardrail_basic(t *testing.T) {
	ctx := acctest.Context(t)

	var guardrail bedrock.GetGuardrailOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_bedrock_guardrail.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGuardrailDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGuardrailConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuardrailExists(ctx, resourceName, &guardrail),
					resource.TestCheckResourceAttrSet(resourceName, "guardrail_arn"),
					resource.TestCheckResourceAttr(resourceName, "blocked_input_messaging", "test"),
					resource.TestCheckResourceAttr(resourceName, "blocked_outputs_messaging", "test"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.0.filters_config.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.0.tier_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "contextual_grounding_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "contextual_grounding_policy_config.0.filters_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cross_region_config.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "test"),
					resource.TestCheckNoResourceAttr(resourceName, names.AttrKMSKeyARN),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.0.pii_entities_config.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.0.regexes_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrStatus, "READY"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.0.topics_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.0.tier_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "DRAFT"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.0.managed_word_lists_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.0.words_config.#", "1"),
				),
			},
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccGuardrailImportStateIDFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "guardrail_id",
			},
		},
	})
}

func TestAccBedrockGuardrail_disappears(t *testing.T) {
	ctx := acctest.Context(t)

	var guardrail bedrock.GetGuardrailOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_bedrock_guardrail.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGuardrailDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGuardrailConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuardrailExists(ctx, resourceName, &guardrail),
					acctest.CheckFrameworkResourceDisappears(ctx, acctest.Provider, tfbedrock.ResourceGuardrail, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccBedrockGuardrail_kmsKey(t *testing.T) {
	ctx := acctest.Context(t)

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_bedrock_guardrail.test"
	var guardrail bedrock.GetGuardrailOutput

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGuardrailDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGuardrailConfig_kmsKey(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuardrailExists(ctx, resourceName, &guardrail),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrKMSKeyARN),
				),
			},
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccGuardrailImportStateIDFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "guardrail_id",
			},
		},
	})
}

func TestAccBedrockGuardrail_update(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_bedrock_guardrail.test"
	var guardrail bedrock.GetGuardrailOutput

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGuardrailDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGuardrailConfig_wordConfig_only(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuardrailExists(ctx, resourceName, &guardrail),
					resource.TestCheckResourceAttrSet(resourceName, "guardrail_arn"),
					resource.TestCheckResourceAttr(resourceName, "blocked_input_messaging", "test"),
					resource.TestCheckResourceAttr(resourceName, "blocked_outputs_messaging", "test"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrCreatedAt),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "test"),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrStatus, "READY"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.#", "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrVersion, "DRAFT"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.0.managed_word_lists_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.0.words_config.#", "1"),
				),
			},
			{
				Config: testAccGuardrailConfig_update(rName, "test", "test", "MEDIUM", "^\\d{3}-\\d{2}-\\d{4}$", "NAME", "investment_topic", "HATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuardrailExists(ctx, resourceName, &guardrail),
					resource.TestCheckResourceAttr(resourceName, "blocked_input_messaging", "test"),
					resource.TestCheckResourceAttr(resourceName, "blocked_outputs_messaging", "test"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.0.filters_config.0.input_strength", "MEDIUM"),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.0.regexes_config.0.pattern", "^\\d{3}-\\d{2}-\\d{4}$"),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.0.pii_entities_config.0.type", "NAME"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.0.topics_config.0.name", "investment_topic"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.0.words_config.0.text", "HATE"),
				),
			},
			{
				Config: testAccGuardrailConfig_update(rName, "update", "update", "HIGH", "^\\d{4}-\\d{2}-\\d{4}$", "USERNAME", "earnings_topic", "HATRED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuardrailExists(ctx, resourceName, &guardrail),
					resource.TestCheckResourceAttr(resourceName, "blocked_input_messaging", "update"),
					resource.TestCheckResourceAttr(resourceName, "blocked_outputs_messaging", "update"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.0.filters_config.0.input_strength", "HIGH"),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.0.regexes_config.0.pattern", "^\\d{4}-\\d{2}-\\d{4}$"),
					resource.TestCheckResourceAttr(resourceName, "sensitive_information_policy_config.0.pii_entities_config.0.type", "USERNAME"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.0.topics_config.0.name", "earnings_topic"),
					resource.TestCheckResourceAttr(resourceName, "word_policy_config.0.words_config.0.text", "HATRED"),
				),
			},
		},
	})
}

func TestAccBedrockGuardrail_crossRegion(t *testing.T) {
	ctx := acctest.Context(t)

	var guardrail bedrock.GetGuardrailOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_bedrock_guardrail.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.BedrockEndpointID)
			acctest.PreCheckRegion(t, endpoints.UsWest2RegionID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.BedrockServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGuardrailDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGuardrailConfig_crossRegion(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGuardrailExists(ctx, resourceName, &guardrail),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.0.tier_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "content_policy_config.0.tier_config.0.tier_name", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "cross_region_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.0.tier_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "topic_policy_config.0.tier_config.0.tier_name", "STANDARD"),
				),
			},
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccGuardrailImportStateIDFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "guardrail_id",
			},
		},
	})
}

func testAccGuardrailImportStateIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("%s,%s", rs.Primary.Attributes["guardrail_id"], rs.Primary.Attributes[names.AttrVersion]), nil
	}
}

func testAccCheckGuardrailDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).BedrockClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_bedrock_guardrail" {
				continue
			}

			id := rs.Primary.Attributes["guardrail_id"]
			version := rs.Primary.Attributes[names.AttrVersion]

			_, err := tfbedrock.FindGuardrailByTwoPartKey(ctx, conn, id, version)
			if errs.IsA[*types.ResourceNotFoundException](err) {
				return nil
			}
			if err != nil {
				return create.Error(names.Bedrock, create.ErrActionCheckingDestroyed, tfbedrock.ResNameGuardrail, rs.Primary.ID, err)
			}

			return create.Error(names.Bedrock, create.ErrActionCheckingDestroyed, tfbedrock.ResNameGuardrail, rs.Primary.ID, errors.New("not destroyed"))
		}

		return nil
	}
}

func testAccCheckGuardrailExists(ctx context.Context, name string, guardrail *bedrock.GetGuardrailOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return create.Error(names.Bedrock, create.ErrActionCheckingExistence, tfbedrock.ResNameGuardrail, name, errors.New("not found"))
		}

		id := rs.Primary.Attributes["guardrail_id"]
		version := rs.Primary.Attributes[names.AttrVersion]
		if id == "" {
			return create.Error(names.Bedrock, create.ErrActionCheckingExistence, tfbedrock.ResNameGuardrail, name, errors.New("guardrail_id not set"))
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).BedrockClient(ctx)

		out, err := tfbedrock.FindGuardrailByTwoPartKey(ctx, conn, id, version)
		if err != nil {
			return create.Error(names.Bedrock, create.ErrActionCheckingExistence, tfbedrock.ResNameGuardrail, rs.Primary.ID, err)
		}

		*guardrail = *out

		return nil
	}
}

func testAccGuardrailConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_bedrock_guardrail" "test" {
  name                      = %[1]q
  blocked_input_messaging   = "test"
  blocked_outputs_messaging = "test"
  description               = "test"

  content_policy_config {
    filters_config {
      input_strength  = "MEDIUM"
      output_strength = "MEDIUM"
      type            = "HATE"
    }
    filters_config {
      input_strength  = "HIGH"
      output_strength = "HIGH"
      type            = "VIOLENCE"
    }
  }

  contextual_grounding_policy_config {
    filters_config {
      threshold = 0.4
      type      = "GROUNDING"
    }
  }

  sensitive_information_policy_config {
    pii_entities_config {
      action = "BLOCK"
      type   = "NAME"
    }
    pii_entities_config {
      action = "BLOCK"
      type   = "DRIVER_ID"
    }
    pii_entities_config {
      action = "ANONYMIZE"
      type   = "USERNAME"
    }
    regexes_config {
      action      = "BLOCK"
      description = "example regex"
      name        = "regex_example"
      pattern     = "^\\d{3}-\\d{2}-\\d{4}$"
    }
  }

  topic_policy_config {
    topics_config {
      name       = "investment_topic"
      examples   = ["Where should I invest my money ?"]
      type       = "DENY"
      definition = "Investment advice refers to inquiries, guidance, or recommendations regarding the management or allocation of funds or assets with the goal of generating returns ."
    }
  }

  word_policy_config {
    managed_word_lists_config {
      type = "PROFANITY"
    }
    words_config {
      text = "HATE"
    }
  }
}
`, rName)
}

func testAccGuardrailConfig_kmsKey(rName string) string {
	return acctest.ConfigCompose(
		testAccCustomModelConfig_base(rName),
		fmt.Sprintf(`
resource "aws_kms_key" "test" {
  description             = %[1]q
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

resource "aws_bedrock_guardrail" "test" {
  name                      = %[1]q
  blocked_input_messaging   = "test"
  blocked_outputs_messaging = "test"
  description               = "test"
  kms_key_arn               = aws_kms_key.test.arn

  content_policy_config {
    filters_config {
      input_strength  = "MEDIUM"
      output_strength = "MEDIUM"
      type            = "HATE"
    }
    filters_config {
      input_strength  = "HIGH"
      output_strength = "HIGH"
      type            = "VIOLENCE"
    }
  }

  word_policy_config {
    managed_word_lists_config {
      type = "PROFANITY"
    }
    words_config {
      text = "HATE"
    }
  }
}
`, rName))
}

func testAccGuardrailConfig_update(rName, blockedInputMessaging, blockedOutputMessaging, inputStrength, regexPattern, piiType, topicName, wordConfig string) string {
	return fmt.Sprintf(`
resource "aws_bedrock_guardrail" "test" {
  name                      = %[1]q
  blocked_input_messaging   = %[2]q
  blocked_outputs_messaging = %[3]q
  description               = "test"

  content_policy_config {
    filters_config {
      input_strength  = %[4]q
      output_strength = "MEDIUM"
      type            = "HATE"
    }
  }

  sensitive_information_policy_config {
    pii_entities_config {
      action = "BLOCK"
      type   = %[6]q
    }
    regexes_config {
      action      = "BLOCK"
      description = "example regex"
      name        = "regex_example"
      pattern     = %[5]q
    }
  }

  topic_policy_config {
    topics_config {
      name       = %[7]q
      examples   = ["Where should I invest my money ?"]
      type       = "DENY"
      definition = "Investment advice refers to inquiries, guidance, or recommendations regarding the management or allocation of funds or assets with the goal of generating returns ."
    }
  }

  word_policy_config {
    managed_word_lists_config {
      type = "PROFANITY"
    }
    words_config {
      text = %[8]q
    }
  }
}
`, rName, blockedInputMessaging, blockedOutputMessaging, inputStrength, regexPattern, piiType, topicName, wordConfig)
}

func testAccGuardrailConfig_wordConfig_only(rName string) string {
	return acctest.ConfigCompose(
		testAccCustomModelConfig_base(rName),
		fmt.Sprintf(`
resource "aws_bedrock_guardrail" "test" {
  name                      = %[1]q
  blocked_input_messaging   = "test"
  blocked_outputs_messaging = "test"
  description               = "test"

  word_policy_config {
    managed_word_lists_config {
      type = "PROFANITY"
    }
    words_config {
      text = "HATE"
    }
  }
}
`, rName))
}

func testAccGuardrailConfig_crossRegion(rName string) string {
	return fmt.Sprintf(`
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_partition" "current" {}

resource "aws_bedrock_guardrail" "test" {
  name                      = %[1]q
  blocked_input_messaging   = "test"
  blocked_outputs_messaging = "test"
  description               = "test"

  cross_region_config {
    guardrail_profile_identifier = "arn:${data.aws_partition.current.partition}:bedrock:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:guardrail-profile/us.guardrail.v1:0"
  }

  content_policy_config {
    filters_config {
      input_strength  = "MEDIUM"
      output_strength = "MEDIUM"
      type            = "HATE"
    }
    filters_config {
      input_strength  = "HIGH"
      output_strength = "HIGH"
      type            = "VIOLENCE"
    }
    tier_config {
      tier_name = "STANDARD"
    }
  }

  contextual_grounding_policy_config {
    filters_config {
      threshold = 0.4
      type      = "GROUNDING"
    }
  }

  sensitive_information_policy_config {
    pii_entities_config {
      action = "BLOCK"
      type   = "NAME"
    }
    pii_entities_config {
      action = "BLOCK"
      type   = "DRIVER_ID"
    }
    pii_entities_config {
      action = "ANONYMIZE"
      type   = "USERNAME"
    }
    regexes_config {
      action      = "BLOCK"
      description = "example regex"
      name        = "regex_example"
      pattern     = "^\\d{3}-\\d{2}-\\d{4}$"
    }
  }

  topic_policy_config {
    topics_config {
      name       = "investment_topic"
      examples   = ["Where should I invest my money ?"]
      type       = "DENY"
      definition = "Investment advice refers to inquiries, guidance, or recommendations regarding the management or allocation of funds or assets with the goal of generating returns ."
    }
    tier_config {
      tier_name = "STANDARD"
    }
  }

  word_policy_config {
    managed_word_lists_config {
      type = "PROFANITY"
    }
    words_config {
      text = "HATE"
    }
  }
}
`, rName)
}
