<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: status.proto

namespace Memphisphp\Pb;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>pb.Status</code>
 */
class Status extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>optional string text = 1;</code>
     */
    protected $text = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $text
     * }
     */
    public function __construct($data = NULL) {
        \Memphisphp\Metadata\Status::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>optional string text = 1;</code>
     * @return string
     */
    public function getText()
    {
        return isset($this->text) ? $this->text : '';
    }

    public function hasText()
    {
        return isset($this->text);
    }

    public function clearText()
    {
        unset($this->text);
    }

    /**
     * Generated from protobuf field <code>optional string text = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setText($var)
    {
        GPBUtil::checkString($var, True);
        $this->text = $var;

        return $this;
    }

}

