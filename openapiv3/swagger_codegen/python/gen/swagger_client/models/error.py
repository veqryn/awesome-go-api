# coding: utf-8

"""
    Awesome GO API

    Actual example use cases for a curated list of golang api generator libraries  # noqa: E501

    OpenAPI spec version: 1.0.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""

import pprint
import re  # noqa: F401

import six

class Error(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """
    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'title': 'str',
        'details': 'str',
        'properties': 'dict(str, object)'
    }

    attribute_map = {
        'title': 'title',
        'details': 'details',
        'properties': 'properties'
    }

    def __init__(self, title=None, details=None, properties=None):  # noqa: E501
        """Error - a model defined in Swagger"""  # noqa: E501
        self._title = None
        self._details = None
        self._properties = None
        self.discriminator = None
        self.title = title
        if details is not None:
            self.details = details
        if properties is not None:
            self.properties = properties

    @property
    def title(self):
        """Gets the title of this Error.  # noqa: E501

        A short, human-readable summary of the problem type. This value should not change between occurrences of the error.  # noqa: E501

        :return: The title of this Error.  # noqa: E501
        :rtype: str
        """
        return self._title

    @title.setter
    def title(self, title):
        """Sets the title of this Error.

        A short, human-readable summary of the problem type. This value should not change between occurrences of the error.  # noqa: E501

        :param title: The title of this Error.  # noqa: E501
        :type: str
        """
        if title is None:
            raise ValueError("Invalid value for `title`, must not be `None`")  # noqa: E501

        self._title = title

    @property
    def details(self):
        """Gets the details of this Error.  # noqa: E501

        A human-readable explanation specific to this occurrence of the problem.  # noqa: E501

        :return: The details of this Error.  # noqa: E501
        :rtype: str
        """
        return self._details

    @details.setter
    def details(self, details):
        """Sets the details of this Error.

        A human-readable explanation specific to this occurrence of the problem.  # noqa: E501

        :param details: The details of this Error.  # noqa: E501
        :type: str
        """

        self._details = details

    @property
    def properties(self):
        """Gets the properties of this Error.  # noqa: E501

        Optional map of properties  # noqa: E501

        :return: The properties of this Error.  # noqa: E501
        :rtype: dict(str, object)
        """
        return self._properties

    @properties.setter
    def properties(self, properties):
        """Sets the properties of this Error.

        Optional map of properties  # noqa: E501

        :param properties: The properties of this Error.  # noqa: E501
        :type: dict(str, object)
        """

        self._properties = properties

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(Error, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, Error):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other