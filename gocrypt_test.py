"""                                                                                                                     
Unit testing for Encryption
"""
import unittest
from gocrypt import Encryption

passphrase = "thisisnotagoodpassphrase"
payload = "plaintext"


class TestCrypto(unittest.TestCase):
    def setUp(self):
        e = Encryption()
        self.encrypted = e.encrypt_string(payload, passphrase)
        self.decrypted = e.decrypt_string(self.encrypted, passphrase)

    def test_encryption(self):
        self.assertEqual(
            self.encrypted, "MFZsSUVMVEk0L1loLSZukGqznR0xu2m4DfSfa+b1HoPVYKC2KQ=="
        )

    def test_decryption(self):
        self.assertEqual(self.decrypted, payload)
