# coding: utf-8

"""
    Awesome GO API

    Actual example use cases for a curated list of golang api generator libraries

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from openapi_client.models.awesome_review_req import AwesomeReviewReq

class TestAwesomeReviewReq(unittest.TestCase):
    """AwesomeReviewReq unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AwesomeReviewReq:
        """Test AwesomeReviewReq
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AwesomeReviewReq`
        """
        model = AwesomeReviewReq()
        if include_optional:
            return AwesomeReviewReq(
                author = '',
                message = '',
                rating = 1
            )
        else:
            return AwesomeReviewReq(
                author = '',
                rating = 1,
        )
        """

    def testAwesomeReviewReq(self):
        """Test AwesomeReviewReq"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()